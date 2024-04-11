package concurrent

import "sync"

type Tasker interface {
	Execute() error
}

type Option func(*Options)

type Options struct {
	ConcurrentNum   int32
	Tasks           []Tasker
	WG              *sync.WaitGroup
	ShowProgressbar bool
}

func NewOptions(opts ...Option) Options {
	opt := Options{
		ConcurrentNum: 1,
		WG:            new(sync.WaitGroup),
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func ConcurrentNum(n int32) Option {
	return func(o *Options) {
		o.ConcurrentNum = n
	}
}

func Tasks(t []Tasker) Option {
	return func(o *Options) {
		o.Tasks = t
	}
}

func WG(wg *sync.WaitGroup) Option {
	return func(o *Options) {
		o.WG = wg
	}
}

func ShowProgressbar() Option {
	return func(o *Options) {
		o.ShowProgressbar = true
	}
}
