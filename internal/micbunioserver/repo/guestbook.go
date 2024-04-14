package repo

import (
	"context"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/model"
	"github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"gorm.io/gorm"
	"time"
)

type Guestbook struct {
	db *gorm.DB
}

func NewGuestbook(db *gorm.DB) *Guestbook {
	return &Guestbook{db: db}
}

func (g *Guestbook) CreateGuestbook(ctx context.Context, req *micbunio.CreateGuestbookRequest) error {
	var guestbookModel = &model.Guestbook{
		Name:    req.Name,
		Content: req.Content,
		HostURL: req.HostUrl,
	}

	return g.db.WithContext(ctx).Create(guestbookModel).Error
}

func (g *Guestbook) GetGuestbookList(ctx context.Context, req *micbunio.GetGuestbookListRequest) (*micbunio.GetGuestbookListResponse, error) {
	var guestbookModels []model.Guestbook
	if err := g.db.WithContext(ctx).
		Where(`host_url = ?`, req.HostUrl).
		Limit(int(req.Limit)).
		Offset(int(req.Offset)).
		Order("created_at desc").
		Find(&guestbookModels).
		Error; err != nil {
		return nil, err
	}

	var guestbooks []*micbunio.Guestbook
	for _, guestbookModel := range guestbookModels {
		guestbooks = append(guestbooks, &micbunio.Guestbook{
			Id:        int64(guestbookModel.ID),
			Name:      guestbookModel.Name,
			Content:   guestbookModel.Content,
			CreatedAt: guestbookModel.Model.CreatedAt.Format(time.RFC3339),
		})
	}

	return &micbunio.GetGuestbookListResponse{
		Guestbooks: guestbooks,
	}, nil
}
