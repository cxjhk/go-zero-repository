package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-repository/user/api/internal/config"
	"go-zero-repository/user/api/internal/middleware"
)

type ServiceContext struct {
	Config        config.Config
	ErrMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ErrMiddleware: middleware.NewErrMiddleware().Handle,
	}
}
