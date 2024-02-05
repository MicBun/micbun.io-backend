package blog

import (
	"context"
	"fmt"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/blog/repo"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/db"
	pb "github.com/MicBun/micbun.io-backend/rpc/micbunio"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"time"
)

type PbBlogServer struct {
	redisManager db.RedisManager
	blogRepo     *repo.Blog
	commentRepo  *repo.Comment
}

func NewPbBlogServer(
	redisManager db.RedisManager,
	blogRepo *repo.Blog,
	commentRepo *repo.Comment,
) *PbBlogServer {
	return &PbBlogServer{
		redisManager: redisManager,
		blogRepo:     blogRepo,
		commentRepo:  commentRepo,
	}
}

func (s *PbBlogServer) ListTest(_ context.Context, _ *pb.GetBlogListRequest) (*pb.GetBlogListResponse, error) {
	return &pb.GetBlogListResponse{
		Blogs: []*pb.Blog{
			{
				Id:        1,
				Title:     "title 1",
				Content:   "content 1",
				CreatedAt: time.Now().Format(time.RFC3339),
			},
			{
				Id:        2,
				Title:     "title 1",
				Content:   "content 2",
				CreatedAt: time.Now().Format(time.RFC3339),
			},
		},
	}, nil
}

func (s *PbBlogServer) GetBlogList(ctx context.Context, in *pb.GetBlogListRequest) (*pb.GetBlogListResponse, error) {
	var (
		blogs *pb.GetBlogListResponse
		err   error
	)

	if err = s.redisManager.Get(ctx, "blogs", &blogs); err != nil {
		return nil, err
	}

	if blogs != nil {
		return blogs, nil
	}

	blogs, err = s.blogRepo.GetBlogList(ctx, in)
	if err != nil {
		return nil, err
	}

	return blogs, s.redisManager.SetEx(ctx, "blogs", blogs, 5*time.Minute)
}

func (s *PbBlogServer) GetBlog(ctx context.Context, in *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	var (
		blog *pb.GetBlogResponse
		err  error
	)

	if err = s.redisManager.Get(ctx, fmt.Sprintf("blog-%d", in.Id), &blog); err != nil {
		return nil, err
	}

	blog, err = s.blogRepo.GetBlog(ctx, in)
	if err != nil {
		return nil, err
	}

	return blog, s.redisManager.SetEx(ctx, "blog-"+fmt.Sprint(in.Id), blog, 5*time.Minute)
}

func (s *PbBlogServer) CreateComment(ctx context.Context, in *pb.CreateCommentRequest) (*google_protobuf.Empty, error) {
	if err := s.commentRepo.CreateComment(ctx, in); err != nil {
		return nil, err
	}

	return &google_protobuf.Empty{}, s.redisManager.Delete(ctx, fmt.Sprintf("comments-%d", in.BlogId))
}

func (s *PbBlogServer) GetCommentList(ctx context.Context, in *pb.GetCommentListRequest) (*pb.GetCommentListResponse, error) {
	var (
		comments *pb.GetCommentListResponse
		err      error
	)

	if err = s.redisManager.Get(ctx, fmt.Sprintf("comments-%d", in.BlogId), &comments); err != nil {
		fmt.Println("redis error", err)
		return nil, err
	}
	if comments != nil {
		return comments, nil
	}

	comments, err = s.commentRepo.GetCommentList(ctx, in)
	if err != nil {
		return nil, err
	}

	return comments, s.redisManager.SetEx(ctx, fmt.Sprintf("comments-%d", in.BlogId), comments, 5*time.Minute)
}
