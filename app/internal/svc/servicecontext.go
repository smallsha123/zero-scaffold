package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-scaffold/app/internal/config"
	"zero-scaffold/model"
)

type ServiceContext struct {
	Config          config.Config
	Redis           *redis.Redis
	WxOfficeAccount model.WxOfficeAccountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		Redis: redis.New(c.RedisConf.Host, func(r *redis.Redis) {
			r.Type = c.RedisConf.Type
			r.Pass = c.RedisConf.Pass
		}),
		WxOfficeAccount: model.NewWxOfficeAccountModel(conn, c.CacheRedis),
	}
}
