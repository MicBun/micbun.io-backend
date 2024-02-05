package repo

import (
	"context"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/blog/model"
	pb "github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	db *gorm.DB
}

func NewBlog(db *gorm.DB) *Blog {
	return &Blog{db: db}
}

func (b *Blog) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	blogModel := model.Blog{
		Model: gorm.Model{
			ID: uint(req.Id),
		},
	}

	if err := b.db.WithContext(ctx).First(&blogModel).Error; err != nil {
		return nil, err
	}

	return &pb.GetBlogResponse{
		Blog: &pb.Blog{
			Id:        int64(blogModel.ID),
			Title:     blogModel.Title,
			Content:   blogModel.Content,
			CreatedAt: blogModel.Model.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (b *Blog) GetBlogList(ctx context.Context, req *pb.GetBlogListRequest) (*pb.GetBlogListResponse, error) {
	var blogModels []model.Blog
	if err := b.db.WithContext(ctx).
		Find(&blogModels).
		Limit(int(req.Limit)).
		Offset(int(req.Offset)).
		Error; err != nil {
		return nil, err
	}

	var blogs []*pb.Blog
	for _, blogModel := range blogModels {
		blogs = append(blogs, &pb.Blog{
			Id:        int64(blogModel.ID),
			Title:     blogModel.Title,
			Content:   blogModel.Content,
			CreatedAt: blogModel.Model.CreatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetBlogListResponse{
		Blogs: blogs,
	}, nil
}
