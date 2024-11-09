package httpc

import (
	"context"
	"github.com/go-resty/resty/v2"
	"sync"
)

// 定义http 引擎
var engine *resty.Client
var once sync.Once

func init() {
	once.Do(func() {
		//TODO 这儿看是否要实现链接池那些 https://github.com/go-resty/resty
		engine = MustNewHttpClient()
	})
}

func Do(ctx context.Context) *resty.Request {
	return engine.R().SetContext(ctx)
}

// MustNewHttpClient new http client
func MustNewHttpClient() *resty.Client {
	return resty.New()
}
