package concurrent

import (
	"context"
)

type Tasker interface {
	Execute(ctx context.Context) error
}

type Option func(*Options)

type Options struct {
	MaxG            int
	Tasks           []Tasker
	ShowProgressbar bool
}

func NewOptions(opts ...Option) Options {
	opt := Options{
		MaxG: 1,
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func MaxG(n int) Option {
	return func(o *Options) {
		o.MaxG = n
	}
}

func Tasks(t []Tasker) Option {
	return func(o *Options) {
		o.Tasks = t
	}
}

func ShowProgressbar() Option {
	return func(o *Options) {
		o.ShowProgressbar = true
	}
}
