package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient
}

func CloseRedis() {
	RedisClient.Close()
}
