package pprof

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/hlf513/go-pkg/framework/gin/middleware"
)

func NewHandlerMiddleware(opts ...Option) middleware.HandlerMiddleware {
	opt := newOptions(opts...)
	return func(engine *gin.Engine) {
		pprof.Register(engine, opt.Path)
	}
}
