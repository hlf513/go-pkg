package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Host:        "127.0.0.1",
		Port:        3306,
		Username:    "user",
		Password:    "pwd",
		Database:    "dbname",
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
	Host        string
	Port        int32
	Username    string
	Password    string
	Database    string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
	GormConfig  gorm.Config
}

func GormConfig(c gorm.Config) Option {
	return func(o *Options) {
		o.GormConfig = c
	}
}

func Host(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func Port(port int32) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func Username(username string) Option {
	return func(o *Options) {
		o.Username = username
	}
}

func Password(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}

func Database(database string) Option {
	return func(o *Options) {
		o.Database = database
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
