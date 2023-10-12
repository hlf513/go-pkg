package clickhouse

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
		Name:         "default",
		Host:         "127.0.0.1",
		Port:         3306,
		Username:     "user",
		Password:     "pwd",
		Database:     "dbname",
		ReadTimeout:  10,
		WriteTimeout: 20,
		MaxIdleConn:  10,
		MaxOpenConn:  100,
		MaxLifeTime:  1800 * time.Second,
		LogLevel:     Silent,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	Name         string
	Host         string
	Port         int32
	Username     string
	Password     string
	Database     string
	MaxIdleConn  int
	MaxOpenConn  int
	MaxLifeTime  time.Duration
	ReadTimeout  int
	WriteTimeout int
	LogLevel     string
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func ReadTimeout(t int) Option {
	return func(o *Options) {
		o.ReadTimeout = t
	}
}

func WriteTimeout(t int) Option {
	return func(o *Options) {
		o.WriteTimeout = t
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
