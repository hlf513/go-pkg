package viper

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hlf513/go-pkg/config"
	"github.com/spf13/viper"
	"log"
)

type FileConf struct {
	opts []config.Config
}

// NewFile 文件配置
//
//	conf := viper.NewFile(
//		config.Name("config"),
//		config.Path("$HOME/demo"),
//		config.AddConfig(
//			new(redis.Config),
//			new(zap.Config),
//			new(openObserve.Config),
//			new(mysql.Config),
//		),
//		config.Watcher(false)
//	)
//	defer conf.Close()
func NewFile(opt ...config.Option) FileConf {
	opts := config.NewOptions(opt...)
	viper.SetConfigName(opts.Name)
	viper.SetConfigType(opts.Type)
	for _, path := range opts.Path {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("config not found")
		} else {
			log.Fatalln(err.Error())
		}
	}

	for _, conf := range opts.Configs {
		if err := viper.Unmarshal(conf); err != nil {
			log.Fatalln(err.Error())
		}
		if err := conf.Init(); err != nil {
			log.Fatalln(err.Error())
		}
	}

	if opts.Watcher {
		viper.OnConfigChange(func(in fsnotify.Event) {
			for _, conf := range opts.Configs {
				if err := viper.Unmarshal(conf); err != nil {
					log.Fatalln(err.Error())
				}
				if err := conf.Init(); err != nil {
					log.Fatalln(err.Error())
				}
			}
		})
		viper.WatchConfig()
	}

	return FileConf{opts: opts.Configs}
}

func (f FileConf) Init() error {
	return nil
}

func (f FileConf) Close() error {
	for _, o := range f.opts {
		_ = o.Close()
	}
	return nil
}
