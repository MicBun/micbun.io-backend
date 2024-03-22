package repo

import (
	"context"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/blog/model"
	blog "github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) *Comment {
	return &Comment{db: db}
}

func (c *Comment) GetCommentList(ctx context.Context, req *blog.GetCommentListRequest) (*blog.GetCommentListResponse, error) {
	var commentModels []model.Comment
	if err := c.db.WithContext(ctx).
		Limit(int(req.Limit)).
		Offset(int(req.Offset)).
		Order("created_at desc").
		Find(&commentModels).
		Error; err != nil {
		return nil, err
	}

	var comments []*blog.Comment
	for _, commentModel := range commentModels {
		comments = append(comments, &blog.Comment{
			Id:        int64(commentModel.ID),
			BlogId:    int64(commentModel.BlogID),
			Name:      commentModel.Name,
			Content:   commentModel.Content,
			CreatedAt: commentModel.Model.CreatedAt.Format(time.RFC3339),
		})
	}

	return &blog.GetCommentListResponse{
		Comments: comments,
	}, nil
}

func (c *Comment) CreateComment(ctx context.Context, req *blog.CreateCommentRequest) error {
	return c.db.WithContext(ctx).Debug().Create(&model.Comment{
		BlogID:  uint(req.BlogId),
		Content: req.Content,
		Name:    req.Name,
	}).Error
}
