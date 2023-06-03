package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClientImpl struct {
	client *redis.Client
}

func NewRedisClient(address string) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return &RedisClientImpl{
		client: client,
	}
}

func (rc *RedisClientImpl) Set(ctx context.Context, key string, value interface{}) error {
	return rc.client.Set(ctx, key, value, 0).Err()
}
