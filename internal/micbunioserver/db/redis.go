package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

type Redis struct {
	client *redis.Client
}

type RedisManager interface {
	Get(ctx context.Context, key string, value any)
	SetEx(ctx context.Context, key string, value any, duration time.Duration)
	Delete(ctx context.Context, key string)
}

func NewRedis() *redis.Client {
	address := os.Getenv("REDIS_INTERNAL_URL")
	if address == "" {
		address = os.Getenv("REDIS_EXTERNAL_URL")
	}
	if address == "" {
		address = os.Getenv("REDIS_HOST") + ":" + fmt.Sprint(os.Getenv("REDIS_PORT"))
	}
	if address == "" {
		return nil
	}

	opt, err := redis.ParseURL(address)
	if err != nil {
		log.Println("redis parse url error: ", err)
	}

	return redis.NewClient(opt)
}

func NewRedisManager(client *redis.Client) *Redis {
	return &Redis{client}
}

func (m *Redis) Get(ctx context.Context, key string, value any) {
	result, err := m.client.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Println("redis get error: ", errors.WithStack(err))
		return
	}

	if err = json.Unmarshal([]byte(result), value); err != nil {
		log.Println("redis unmarshal error: ", errors.WithStack(err))
	}
}

func (m *Redis) SetEx(ctx context.Context, key string, value any, duration time.Duration) {
	rawValue, err := json.Marshal(value)
	if err != nil {
		log.Println("redis marshal error: ", errors.WithStack(err))
		return
	}

	if err = m.client.Set(ctx, key, rawValue, duration).Err(); err != nil {
		log.Println("redis set error: ", errors.WithStack(err))
	}
}

func (m *Redis) Delete(ctx context.Context, key string) {
	if err := m.client.Del(ctx, key).Err(); err != nil {
		log.Println("redis delete error: ", errors.WithStack(err))
	}
}
