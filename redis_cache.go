package main

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{
		client: c,
	}
}

func (c *RedisCache) Get(key int) (string, error) {
	ctx := context.Background()
	intStr := strconv.Itoa(key)
	return c.client.Get(ctx, intStr).Result()
}
