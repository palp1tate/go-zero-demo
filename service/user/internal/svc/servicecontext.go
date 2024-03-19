package svc

import (
	"github.com/palp1tate/go-zero-demo/internal/model"
	"github.com/palp1tate/go-zero-demo/pkg/orm"
	"github.com/palp1tate/go-zero-demo/service/user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		Host:     c.Mysql.Host,
		Port:     c.Mysql.Port,
		User:     c.Mysql.User,
		Password: c.Mysql.Password,
		Database: c.Mysql.Database,
	})
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(db),
	}
}
