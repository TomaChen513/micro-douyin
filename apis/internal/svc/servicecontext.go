package svc

import (
	"miniTikTok/apis/internal/config"
	"miniTikTok/apis/internal/middleware"
	"miniTikTok/rpc/user/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
	UserService userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,
		UserService: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
