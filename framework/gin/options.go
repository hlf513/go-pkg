package gin

import (
	"github.com/hlf513/go-pkg/config"
	"github.com/hlf513/go-pkg/framework/gin/middleware"
	"github.com/hlf513/go-pkg/framework/gin/middleware/zap"
)

type Option func(*Options)

type Options struct {
	Middleware []middleware.HandlerMiddleware
	Config     config.Config
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Middleware: []middleware.HandlerMiddleware{
			zap.NewHandlerMiddleware(),
		},
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Middleware(mdl ...middleware.HandlerMiddleware) Option {
	return func(o *Options) {
		o.Middleware = append(o.Middleware, mdl...)
	}
}

func Config(conf config.Config) Option {
	return func(o *Options) {
		o.Config = conf
	}
}
