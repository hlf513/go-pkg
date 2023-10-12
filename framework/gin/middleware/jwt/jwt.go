package jwt

import (
	jwt "github.com/appleboy/gin-jwt/v2"
)

func New(opts ...Option) (*jwt.GinJWTMiddleware, error) {
	opt := newOptions(opts...)
	return jwt.New(&jwt.GinJWTMiddleware{
		Key:           []byte(opt.Key),
		Timeout:       opt.Timeout,
		MaxRefresh:    opt.MaxRefresh,
		TokenLookup:   opt.TokenLookup, // 依次查询
		Authenticator: opt.Authenticator,
		PayloadFunc: func(data any) jwt.MapClaims {
			return jwt.MapClaims{jwt.IdentityKey: data}
		},
		LoginResponse:   opt.LoginResponse,
		LogoutResponse:  opt.LogoutResponse,
		Unauthorized:    opt.Unauthorized,
		RefreshResponse: opt.RefreshResponse,
		SendCookie:      opt.SendCookie,
		CookieName:      opt.CookieName,
	})
}
