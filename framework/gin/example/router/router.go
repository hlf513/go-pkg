package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hlf513/go-pkg/database/go-redis"
	"github.com/hlf513/go-pkg/framework/gin/response"
	"time"
)

func RegisterHandler(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, "not found")
	})

	engine.GET("/redis", func(ctx *gin.Context) {
		rdb := redis.GetClient()
		rdb.Set(ctx.Request.Context(), "key", "redis-value", 10*time.Second)
		v, err := rdb.Get(ctx.Request.Context(), "key").Result()
		if err != nil {
			response.Error(ctx, err.Error())
			return
		}
		response.Success(ctx, v)
	})

	// 权限路由
	admin := engine.Group("/admin")
	auth(admin)

	// 以下接口需要授权后才能访问
	admin.GET("/config", func(c *gin.Context) {
		c.JSON(200, "config")
	})
}
