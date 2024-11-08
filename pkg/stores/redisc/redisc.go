package redisc

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewEngine(c Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
	})
	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}
	return rdb
}
