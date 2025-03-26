package services

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)


type TokenBucketRateLimiter struct {
	client    *redis.Client
	bucketKey string
	capacity  int     
	rate      float64
}

func NewRateLimiter(client *redis.Client, bucketKey string, capacity int, rate float64) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		client:    client,
		bucketKey: bucketKey,
		capacity:  capacity,
		rate:      rate,
	}
}

func (rl *TokenBucketRateLimiter) AllowRequest(ctx context.Context) (bool, error) {
	now := time.Now().Unix()
	redisKey := "rate_limit:" + rl.bucketKey

	tokens, err := rl.client.Get(ctx, redisKey).Int()
	if err == redis.Nil {
		tokens = rl.capacity
	} else if err != nil {
		return false, err
	}

	if tokens <= 0 {
		return false, nil
	}

	_, err = rl.client.Decr(ctx, redisKey).Result()
	if err != nil {
		return false, err
	}

	rl.client.Expire(ctx, redisKey, time.Duration(1/rl.rate)*time.Second)

	return true, nil
}
