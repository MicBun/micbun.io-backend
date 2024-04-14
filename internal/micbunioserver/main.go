package micbunioserver

import (
	"context"
	"fmt"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/db"
	"github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"github.com/golang/protobuf/ptypes/empty"
)

type PbService struct {
	redisManager  db.RedisManager
	guestbookRepo GuestbookRepo
}

func NewPbService(
	redisManager db.RedisManager,
	guestbookRepo GuestbookRepo,
) *PbService {
	return &PbService{
		redisManager:  redisManager,
		guestbookRepo: guestbookRepo,
	}
}

func (s *PbService) GetGuestbookList(ctx context.Context, in *micbunio.GetGuestbookListRequest) (*micbunio.GetGuestbookListResponse, error) {
	var (
		guestbooks *micbunio.GetGuestbookListResponse
		err        error
	)

	redisKey := fmt.Sprintf("guestbooks-%s", in.HostUrl)
	s.redisManager.Get(ctx, redisKey, &guestbooks)

	if guestbooks != nil {
		return guestbooks, nil
	}

	guestbooks, err = s.guestbookRepo.GetGuestbookList(ctx, in)
	if err != nil {
		return nil, err
	}

	s.redisManager.SetEx(ctx, redisKey, guestbooks, 5)
	return guestbooks, nil
}

func (s *PbService) CreateGuestbook(ctx context.Context, in *micbunio.CreateGuestbookRequest) (*empty.Empty, error) {
	if err := s.guestbookRepo.CreateGuestbook(ctx, in); err != nil {
		return nil, err
	}

	s.redisManager.Delete(ctx, fmt.Sprintf("guestbooks-%s", in.HostUrl))
	return &empty.Empty{}, nil
}
