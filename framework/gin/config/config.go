package app

import (
	"github.com/gin-gonic/gin"
)

var config Options

type Config struct {
	Gin Options `json:"gin"`
}

func (a Config) Init() error {
	config = NewOptions(
		Name(a.Gin.Name),
		Mode(a.Gin.Mode),
		Port(a.Gin.Port),
		Env(a.Gin.Env),
	)
	if a.Gin.Mode == gin.DebugMode ||
		a.Gin.Mode == gin.ReleaseMode ||
		a.Gin.Mode == gin.TestMode {
		gin.SetMode(a.Gin.Mode)
	}
	return nil
}

func (a Config) Close() error {
	return nil
}

func GetConfig() Options {
	return config
}
