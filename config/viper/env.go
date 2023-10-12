package viper

import (
	"github.com/spf13/viper"
)

func NewEnv() {
	viper.AutomaticEnv()
	//viper.Get("MySQL")
}
