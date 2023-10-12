package casbin

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	response2 "github.com/hlf513/go-pkg/framework/gin/response"
	"github.com/hlf513/go-pkg/log/zap"
	"github.com/spf13/cast"
	"net/http"
)

// CheckPermission 根据不同的错误返回不同的 http code
// - 未找到用户，403
// - casbin 报错，500
// - 没有权限，401
func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(UserIDCtxName)
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		uid := cast.ToInt64(data)
		if uid == SuperUserID {
			c.Next()
			return
		}

		if pass, err := Instance().CheckPermissionForUser(uid, c.Request.URL.Path, c.Request.Method); err != nil {
			zap.Error(c.Request.Context(), err.Error(), "casbin")
			response2.Error(c, ErrCasbin.Error(), response2.ErrorCode(http.StatusInternalServerError))
			c.Abort()
			return
		} else if !pass {
			response2.Error(c, jwt.ErrForbidden.Error(), response2.ErrorCode(http.StatusUnauthorized))
			c.Abort()
			return
		}

		c.Next()
	}
}
