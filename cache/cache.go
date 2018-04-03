package cache

import (
	"github.com/go-redis/redis"
)

var (
	user = redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			PoolSize: 20,
			DB:       0,
		},
	)

	order = redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			PoolSize: 20,
			DB:       1,
		},
	)
)

func Use(name string) *redis.Client {

	switch name {
	case "user":
		return user
	default:
		return order
	}
}
