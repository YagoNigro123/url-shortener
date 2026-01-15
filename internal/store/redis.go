package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(address string, password string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return &RedisClient{client: rdb}
}

func (r *RedisClient) Save(id string, original string) error {
	ctx := context.Background()

	err := r.client.Set(ctx, id, original, 24*time.Hour).Err()

	if err != nil {
		return fmt.Errorf("redis set error: %w", err)
	}

	return nil
}

func (r *RedisClient) Get(id string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, id).Result()

	if err != nil {
		return "", fmt.Errorf("Link not found in cache")
	} else if err != nil {
		return "", fmt.Errorf("Redis get error: %w", err)
	}

	return val, nil
}
