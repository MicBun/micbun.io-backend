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
	Get(ctx context.Context, key string, value any) error
	SetEx(ctx context.Context, key string, value any, duration time.Duration) error
	Delete(ctx context.Context, key string) error
}

func NewRedis() *redis.Client {
	address := os.Getenv("REDIS_INTERNAL_URL")
	if address == "" {
		address = os.Getenv("REDIS_EXTERNAL_URL")
	}
	if address == "" {
		address = os.Getenv("REDIS_HOST") + ":" + fmt.Sprint(os.Getenv("REDIS_PORT"))
	}
	log.Printf("redis address: %s", address)

	opt, err := redis.ParseURL(address)
	if err != nil {
		log.Println("redis parse url error: ", err)
	}

	return redis.NewClient(opt)
}

func NewRedisManager(client *redis.Client) *Redis {
	return &Redis{client}
}

func (m *Redis) Get(ctx context.Context, key string, value any) error {
	result, err := m.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil
	} else if err != nil {
		return errors.WithStack(err)
	}
	log.Printf("get value from redis key: %s, value: %s", key, result)

	return errors.WithStack(json.Unmarshal([]byte(result), value))
}

func (m *Redis) SetEx(ctx context.Context, key string, value any, duration time.Duration) error {
	rawValue, err := json.Marshal(value)
	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(m.client.Set(ctx, key, rawValue, duration).Err())
}

func (m *Redis) Delete(ctx context.Context, key string) error {
	return errors.WithStack(m.client.Del(ctx, key).Err())
}
