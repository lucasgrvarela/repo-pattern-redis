package main

import "context"

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}) error
}
