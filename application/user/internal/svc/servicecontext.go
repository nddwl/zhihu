package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zhihu/application/user/internal/config"
	"zhihu/application/user/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	BizRedis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheConf),
		BizRedis:  redis.MustNewRedis(c.BizRedis),
	}
}
