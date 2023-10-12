package casbin

import (
	"github.com/hlf513/go-pkg/database/go-redis"
	"github.com/hlf513/go-pkg/database/gorm/mysql"
)

type Option func(*Options)

type Options struct {
	Model        string
	DBOptions    mysql.Options
	RedisOptions redis.Options
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Model: defaultModel,
		DBOptions: mysql.Options{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "123123",
			Database: "casbin",
		},
		RedisOptions: redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   0,
		},
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Model(model string) Option {
	return func(o *Options) {
		o.Model = model
	}
}

func DBOptions(opts mysql.Options) Option {
	return func(o *Options) {
		o.DBOptions = opts
	}
}

func RedisOptions(opts redis.Options) Option {
	return func(o *Options) {
		o.RedisOptions = opts
	}
}
