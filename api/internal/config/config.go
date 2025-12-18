package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	MySQL MySQLConfig
	Redis redis.RedisConf
	Auth  AuthConfig
}

type MySQLConfig struct {
	DataSource   string
	MaxIdleConns int
	MaxOpenConns int
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
