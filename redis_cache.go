package main

import "github.com/redis/go-redis/v9"

type RedisCache struct {
	client *redis.Client
}
