package repo_test

import (
	"context"
	"testing"

	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/model"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/repo"
	"github.com/MicBun/micbun.io-backend/rpc/micbunio"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Guestbook{}))
	return db
}

func TestGuestbook_CreateAndList(t *testing.T) {
	db := setupTestDB(t)
	g := repo.NewGuestbook(db)

	ctx := context.Background()
	// create first entry
	err := g.CreateGuestbook(ctx, &micbunio.CreateGuestbookRequest{
		Name:    "John",
		Content: "hi",
		HostUrl: "localhost",
	})
	require.NoError(t, err)

	// create second entry
	err = g.CreateGuestbook(ctx, &micbunio.CreateGuestbookRequest{
		Name:    "Jane",
		Content: "hello",
		HostUrl: "localhost",
	})
	require.NoError(t, err)

	resp, err := g.GetGuestbookList(ctx, &micbunio.GetGuestbookListRequest{
		HostUrl: "localhost",
		Limit:   10,
		Offset:  0,
	})
	require.NoError(t, err)
	require.Len(t, resp.Guestbooks, 2)
	// order should be latest first
	require.Equal(t, "Jane", resp.Guestbooks[0].Name)
	require.Equal(t, "John", resp.Guestbooks[1].Name)

	// dropping table should cause error
	require.NoError(t, db.Migrator().DropTable(&model.Guestbook{}))
	_, err = g.GetGuestbookList(ctx, &micbunio.GetGuestbookListRequest{HostUrl: "localhost"})
	require.Error(t, err)
}

func TestGuestbook_CreateGuestbook_Error(t *testing.T) {
	db := setupTestDB(t)
	g := repo.NewGuestbook(db)
	ctx := context.Background()

	// drop table to force insert error
	require.NoError(t, db.Migrator().DropTable(&model.Guestbook{}))
	err := g.CreateGuestbook(ctx, &micbunio.CreateGuestbookRequest{HostUrl: "localhost"})
	require.Error(t, err)
}
