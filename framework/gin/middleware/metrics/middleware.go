package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/hlf513/go-pkg/framework/gin/middleware"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func NewHandlerMiddleware(opts ...Option) middleware.HandlerMiddleware {
	opt := newOptions(opts...)
	return func(engine *gin.Engine) {
		m := ginmetrics.GetMonitor()
		m.SetMetricPath(opt.Path)
		m.Use(engine)
	}
}
