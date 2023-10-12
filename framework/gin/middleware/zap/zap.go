package zap

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/hlf513/go-pkg/framework/gin/middleware"
	"github.com/hlf513/go-pkg/log/zap"
)

func NewHandlerMiddleware() middleware.HandlerMiddleware {
	return func(engine *gin.Engine) {
		engine.Use(ginzap.RecoveryWithZap(zap.Logger(), true))
	}
}
