package sqlite

import (
	"time"
)

const (
	Silent = "silent"
	Warn   = "warn"
	Error  = "error"
	Info   = "info"
)

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Name:        "default",
		DSN:         "gorm.db",
		MaxIdleConn: 10,
		MaxOpenConn: 100,
		MaxLifeTime: 1800 * time.Second,
		LogLevel:    Silent,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	Name        string
	DSN         string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
	LogLevel    string
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func DSN(dsn string) Option {
	return func(o *Options) {
		o.DSN = dsn
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
