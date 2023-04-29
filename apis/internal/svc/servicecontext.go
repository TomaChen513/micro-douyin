package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"miniTikTok/apis/internal/config"
	"miniTikTok/apis/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
