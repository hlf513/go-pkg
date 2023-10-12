package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hlf513/go-pkg/database/go-redis"
	"github.com/hlf513/go-pkg/database/gorm/mysql"
	"github.com/hlf513/go-pkg/framework/gin/middleware/casbin"
	"github.com/hlf513/go-pkg/framework/gin/middleware/jwt"
	"github.com/hlf513/go-pkg/framework/gin/response"
	"log"
	"time"
)

func auth(admin *gin.RouterGroup) {
	// 账号登录
	authMiddleware, err := jwt.New(
		jwt.Key("custom_jwt_key"),
		jwt.Authenticator(func(c *gin.Context) (any, error) {
			// 登录逻辑，返回值会保存在 jwt 中
			return nil, nil
		}),
		jwt.CookieName(""),
		jwt.Timeout(12*time.Hour),
		jwt.MaxRefresh(12*time.Hour),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	admin.GET("/user/refresh_token", authMiddleware.RefreshHandler)
	admin.POST("/user/login", authMiddleware.LoginHandler)
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		// 登录用户皆可访问的地址
		admin.POST("/user/logout", func(c *gin.Context) {
			response.Success(c, "logout")
		})
		admin.GET("/user/info", func(c *gin.Context) {
			response.Success(c, "info")
		})
	}

	// 权限校验
	if err = casbin.Instance(
		casbin.DBOptions(mysql.GetConfig()),
		casbin.RedisOptions(redis.GetConfig()),
	).RedisWatcher("admin"); err != nil {
		log.Fatal(err.Error())
	}
	// todo 系统首次启动时需要初始化权限表
	//if err = auth.InitRolePermissions(); err != nil {
	//	log.Fatal(err.Error())
	//}
	admin.Use(casbin.CheckPermission())
}
