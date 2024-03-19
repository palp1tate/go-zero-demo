package config

import (
	"github.com/palp1tate/go-zero-demo/pkg/orm"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql orm.Config `json:"mysql"`
}
