package cache

import (
	"github.com/go-redis/redis"
)

var (
	client = redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			PoolSize: 20,
			DB:       0,
		},
	)
)

func Use(name string) *redis.Client {

	return client
}
