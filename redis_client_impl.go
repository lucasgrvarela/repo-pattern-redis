package main

import (
	"context"
	"fmt"

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

func (rc *RedisClientImpl) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key '%s' not found", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}
