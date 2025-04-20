package shared

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		ReadTimeout: 5 * time.Second,
	})
}
