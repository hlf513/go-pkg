package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Option func(*Options)

type Options struct {
	Key             string
	Timeout         time.Duration
	MaxRefresh      time.Duration
	TokenLookup     string
	Authenticator   func(c *gin.Context) (any, error)
	LoginResponse   func(c *gin.Context, code int, token string, expire time.Time)
	LogoutResponse  func(c *gin.Context, code int)
	Unauthorized    func(c *gin.Context, code int, message string)
	RefreshResponse func(c *gin.Context, code int, token string, expire time.Time)
	SendCookie      bool
	CookieName      string
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Key:         "this is a sampler key",
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		TokenLookup: "header: Authorization, query: token",
		Authenticator: func(c *gin.Context) (any, error) {
			return map[string]string{"user": "gin"}, nil
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
				"data": map[string]any{
					"token":  token,
					"expire": expire.Unix(),
				},
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  message,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
				"data": map[string]any{
					"token":  token,
					"expire": expire.Unix(),
				},
			})
		},
	}

	for _, o := range opts {
		o(&opt)
	}
	if opt.CookieName != "" {
		opt.SendCookie = true
		opt.TokenLookup += ", cookie: " + opt.CookieName
	}

	return opt
}

func Key(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}
func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}
func MaxRefresh(t time.Duration) Option {
	return func(o *Options) {
		o.MaxRefresh = t
	}
}
func TokenLookup(tokenLookup string) Option {
	return func(o *Options) {
		o.TokenLookup = tokenLookup
	}
}
func Authenticator(f func(c *gin.Context) (any, error)) Option {
	return func(o *Options) {
		o.Authenticator = f
	}
}
func LoginResponse(f func(c *gin.Context, code int, token string, expire time.Time)) Option {
	return func(o *Options) {
		o.LoginResponse = f
	}
}
func LogoutResponse(f func(c *gin.Context, code int)) Option {
	return func(o *Options) {
		o.LogoutResponse = f
	}
}
func Unauthorized(f func(c *gin.Context, code int, message string)) Option {
	return func(o *Options) {
		o.Unauthorized = f
	}
}
func RefreshResponse(f func(c *gin.Context, code int, token string, expire time.Time)) Option {
	return func(o *Options) {
		o.RefreshResponse = f
	}
}
func CookieName(cookieName string) Option {
	return func(o *Options) {
		o.CookieName = cookieName
	}
}
