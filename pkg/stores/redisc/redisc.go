package redisc

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewEngine(c Config) redis.UniversalClient {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      c.Addrs,
		Username:   c.Username,
		Password:   c.Password,
		MasterName: c.MasterName,
	})

	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}

	if c.Debug { // debug
		rdb.AddHook(DebugHook{})
	}
	return rdb
}
