package sqlite

import (
	"gorm.io/gorm"
	"time"
)

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		DSN:         "gorm.db",
		MaxIdleConn: 10,
		MaxOpenConn: 100,
		MaxLifeTime: 1800 * time.Second,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	DSN         string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
	GormConfig  gorm.Config
}

func DSN(dsn string) Option {
	return func(o *Options) {
		o.DSN = dsn
	}
}

func GormConfig(c gorm.Config) Option {
	return func(o *Options) {
		o.GormConfig = c
	}
}

func MaxIdleConn(num int) Option {
	return func(o *Options) {
		o.MaxIdleConn = num
	}
}
func MaxOpenConn(num int) Option {
	return func(o *Options) {
		o.MaxOpenConn = num
	}
}

func MaxLifeTime(t time.Duration) Option {
	return func(o *Options) {
		o.MaxLifeTime = t
	}
}
