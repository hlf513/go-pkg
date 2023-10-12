package pprof

import "github.com/gin-contrib/pprof"

type Option func(*Options)

type Options struct {
	Path string
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Path: pprof.DefaultPrefix,
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Path(path string) Option {
	return func(o *Options) {
		o.Path = path
	}
}
