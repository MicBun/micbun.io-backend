package micbunioserver_test

import (
	"context"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/db"
	"github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mainTest struct {
	service *micbunioserver.PbService
	repo    *micbunioserver.MockGuestbookRepo
	redis   *db.MockRedisManager
}

func newMainTest(t *testing.T) mainTest {
	mockRedis := db.NewMockRedisManager(t)
	mockRepo := micbunioserver.NewMockGuestbookRepo(t)
	return mainTest{
		service: micbunioserver.NewPbService(mockRedis, mockRepo),
		repo:    mockRepo,
		redis:   mockRedis,
	}
}

func TestNewPbService(t *testing.T) {
	instance := newMainTest(t)
	if instance.service == nil {
		t.Error("NewPbService returned nil")
	}
}

func TestPbService_GetGuestbookList(t *testing.T) {
	t.Run("success - it should return guestbooks", func(t *testing.T) {
		instance := newMainTest(t)
		instance.redis.EXPECT().Get(mock.Anything, mock.Anything, mock.Anything)
		instance.repo.EXPECT().GetGuestbookList(mock.Anything, mock.Anything).Return(&micbunio.GetGuestbookListResponse{}, nil)
		instance.redis.EXPECT().SetEx(mock.Anything, mock.Anything, mock.Anything, mock.Anything)
		_, err := instance.service.GetGuestbookList(context.Background(), &micbunio.GetGuestbookListRequest{HostUrl: "localhost"})
		assert.NoError(t, err)
	})

	t.Run("error - it should return an error if GetGuestbookList returns an error", func(t *testing.T) {
		instance := newMainTest(t)
		instance.redis.EXPECT().Get(mock.Anything, mock.Anything, mock.Anything)
		instance.repo.EXPECT().GetGuestbookList(mock.Anything, mock.Anything).Return(nil, assert.AnError)
		_, err := instance.service.GetGuestbookList(context.Background(), &micbunio.GetGuestbookListRequest{HostUrl: "localhost"})
		assert.EqualError(t, err, assert.AnError.Error())
	})
}

func TestPbService_CreateGuestbook(t *testing.T) {
	t.Run("success - it should return nil", func(t *testing.T) {
		instance := newMainTest(t)
		instance.repo.EXPECT().CreateGuestbook(mock.Anything, mock.Anything).Return(nil)
		instance.redis.EXPECT().Delete(mock.Anything, mock.Anything)
		_, err := instance.service.CreateGuestbook(context.Background(), &micbunio.CreateGuestbookRequest{HostUrl: "localhost"})
		assert.NoError(t, err)
	})

	t.Run("error - it should return an error if CreateGuestbook returns an error", func(t *testing.T) {
		instance := newMainTest(t)
		instance.repo.EXPECT().CreateGuestbook(mock.Anything, mock.Anything).Return(assert.AnError)
		_, err := instance.service.CreateGuestbook(context.Background(), &micbunio.CreateGuestbookRequest{HostUrl: "localhost"})
		assert.EqualError(t, err, assert.AnError.Error())
	})
}
