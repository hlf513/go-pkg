package redis

import "time"

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Name:            "default",
		Server:          "127.0.0.1",
		Port:            "6379",
		Password:        "",
		MaxIdleConn:     1,
		MaxActiveConn:   10,
		ConnectTimeout:  500 * time.Microsecond,
		ReadTimeout:     300 * time.Microsecond,
		WriteTimeout:    300 * time.Microsecond,
		IdleTimeout:     300 * time.Second,
		MaxConnLifetime: 1800 * time.Second,
		SelectDB:        0,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	Name            string
	Server          string
	Port            string
	Password        string
	MaxIdleConn     int
	MaxActiveConn   int
	ConnectTimeout  time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
	SelectDB        int
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func Server(server string) Option {
	return func(o *Options) {
		o.Server = server
	}
}
func Port(port string) Option {
	return func(o *Options) {
		o.Port = port
	}
}
func Password(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}
func MaxIdleConn(num int) Option {
	return func(o *Options) {
		o.MaxIdleConn = num
	}
}
func MaxActiveConn(num int) Option {
	return func(o *Options) {
		o.MaxActiveConn = num
	}
}
func ConnectTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.ConnectTimeout = timeout
	}
}
func ReadTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = timeout
	}
}
func WriteTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = timeout
	}
}
func IdleTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = timeout
	}
}
func MaxConnLifetime(timeout time.Duration) Option {
	return func(o *Options) {
		o.MaxConnLifetime = timeout
	}
}
func SelectDB(num int) Option {
	return func(o *Options) {
		o.SelectDB = num
	}
}
