package db_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"

	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/db"
)

func TestNewRedis(t *testing.T) {
	mr := miniredis.RunT(t)
	os.Setenv("REDIS_INTERNAL_URL", mr.Addr())
	defer os.Unsetenv("REDIS_INTERNAL_URL")

	client := db.NewRedis()
	require.NotNil(t, client)
	require.NoError(t, client.Ping(context.Background()).Err())
}

func TestNewRedis_NoEnv(t *testing.T) {
	os.Unsetenv("REDIS_INTERNAL_URL")
	os.Unsetenv("REDIS_EXTERNAL_URL")
	os.Unsetenv("REDIS_HOST")
	os.Unsetenv("REDIS_PORT")
	require.Nil(t, db.NewRedis())
}

func TestRedisManager(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	mgr := db.NewRedisManager(rdb)
	ctx := context.Background()

	type sample struct{ A string }
	val := sample{A: "B"}

	mgr.SetEx(ctx, "key", val, time.Hour)

	var got sample
	mgr.Get(ctx, "key", &got)
	require.Equal(t, val, got)

	// put invalid json
	rdb.Set(ctx, "bad", "notjson", 0)
	mgr.Get(ctx, "bad", &got) // should not panic and leave got unchanged
	require.Equal(t, val, got)

	mgr.Delete(ctx, "key")
	_, err := rdb.Get(ctx, "key").Result()
	require.Equal(t, redis.Nil, err)
}
