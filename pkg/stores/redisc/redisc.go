package redisc

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewEngine(c Config) redis.UniversalClient {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: c.Addrs,
	})
	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}
	return rdb
}
