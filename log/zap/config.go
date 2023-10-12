package zap

import "log"

var config Options

type Config struct {
	Zap Options `json:"zap"`
}

func (c Config) Init() error {
	config = c.Zap
	Setup(
		LogPath(c.Zap.LogPath),
		FileMaxSize(c.Zap.FileMaxSize),
		FileMaxBackups(c.Zap.FileMaxBackups),
		FileMaxAge(c.Zap.FileMaxAge),
		FileCompress(c.Zap.FileCompress),
		LogLevel(c.Zap.LogLevel),
	)
	log.Print("[Zap] init Zap configure")
	return nil
}

func (c Config) Close() error {
	return nil
}

func GetConfig() Options {
	return config
}
