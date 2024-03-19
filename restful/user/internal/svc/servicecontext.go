package svc

import (
	"github.com/palp1tate/go-zero-demo/restful/user/internal/config"
	"github.com/palp1tate/go-zero-demo/service/user/usersrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usersrv.UserSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usersrv.NewUserSrv(zrpc.MustNewClient(c.UserRpc)),
	}
}
