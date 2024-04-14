package micbunioserver

import (
	"context"
	"github.com/MicBun/micbun.io-backend/rpc/micbunio"
)

type GuestbookRepo interface {
	GetGuestbookList(ctx context.Context, in *micbunio.GetGuestbookListRequest) (*micbunio.GetGuestbookListResponse, error)
	CreateGuestbook(ctx context.Context, in *micbunio.CreateGuestbookRequest) error
}
