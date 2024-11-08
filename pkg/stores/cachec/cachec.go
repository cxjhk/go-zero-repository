package cachec

import (
	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
)

func NewEngine(c Config, rdb *redis.Client) *rockscache.Client {
	options := rockscache.NewDefaultOptions()
	options.StrongConsistency = c.StrongConsistency
	options.DisableCacheRead = c.DisableCacheRead
	return rockscache.NewClient(rdb, options)
}
