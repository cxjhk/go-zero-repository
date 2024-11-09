package redisc

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type DebugHook struct {
	redis.Hook
}

func (DebugHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		logx.Debugf("redis cmd: %s", cmd.String())
		next(ctx, cmd)
		return nil
	}
}
