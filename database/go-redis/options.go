package redis

import (
	"net"
	"time"
)

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Name: "default",
		Addr: net.JoinHostPort("127.0.0.1", "6379"),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	Name         string `json:"name" mapstructure:"name"`
	Addr         string `json:"addr" mapstructure:"addr"`
	Username     string `json:"username" mapstructure:"username"`
	Password     string `json:"password" mapstructure:"password"`
	DB           int    `json:"db" mapstructure:"db"`
	MaxRetries   int    `json:"max_retries" mapstructure:"max_retries"`
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
	Trace        bool
	Metrics      bool
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func Addr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
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

func DB(num int) Option {
	return func(o *Options) {
		o.DB = num
	}
}

func MaxRetries(max int) Option {
	return func(o *Options) {
		if max > 0 {
			o.MaxRetries = max
		}
	}
}

func DialTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		if timeout > 0 {
			o.DialTimeout = timeout
		}
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		if timeout > 0 {
			o.ReadTimeout = timeout
		}
	}
}
func WriteTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		if timeout > 0 {
			o.WriteTimeout = timeout
		}
	}
}

func PoolSize(size int) Option {
	return func(o *Options) {
		if size > 0 {
			o.PoolSize = size
		}
	}
}

func PoolTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		if timeout > 0 {
			o.PoolTimeout = timeout
		}
	}
}

func Trace(b bool) Option {
	return func(o *Options) {
		o.Trace = b
	}
}

func Metrics(b bool) Option {
	return func(o *Options) {
		o.Metrics = b
	}
}
