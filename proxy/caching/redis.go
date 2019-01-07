package caching

import (
	"github.com/go-redis/redis"
)

const (
	HOST = "redis"
	PORT = "6379"
)

var RedisClient *redis.Client

func InitRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     HOST + ":" + PORT,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return RedisClient
}
