package concurrency

import "sync"

type Tasker interface {
	Run() error
}

type Option func(*Options)

type Options struct {
	Concurrency int32
	Tasks       []Tasker
	WG          *sync.WaitGroup
}

func NewOptions(opts ...Option) Options {
	opt := Options{
		Concurrency: 1,
		WG:          new(sync.WaitGroup),
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Concurrency(n int32) Option {
	return func(o *Options) {
		o.Concurrency = n
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
