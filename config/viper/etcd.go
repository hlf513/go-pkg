package viper

import (
	"github.com/hlf513/go-pkg/config"
	"github.com/spf13/viper"
	"log"
)

type EtcdConf struct {
	opts []config.Config
}

// NewEtcd Etcd 配置
//
// conf := viper.NewFile(
//
//	config.Endpoint("http://127.0.0.1:4001"),
//	config.Path("/config/demo.json"),
//	config.AddConfig(
//		new(redis.Config),
//		new(zap.Config),
//		new(openObserve.Config),
//		new(mysql.Config),
//	),
//	config.Watcher(false), // true 开启监控，配置变更执行 Init()
//
// )
// defer conf.Close()
func NewEtcd(opt ...config.Option) EtcdConf {
	opts := config.NewOptions(opt...)
	var runtimeViper = viper.New()
	if err := runtimeViper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001", "/config/hugo.json"); err != nil {
		log.Fatal(err.Error())
	}
	runtimeViper.SetConfigType("json")

	if err := runtimeViper.ReadRemoteConfig(); err != nil {
		log.Fatalln(err.Error())
	}

	for _, conf := range opts.Configs {
		if err := runtimeViper.Unmarshal(conf); err != nil {
			log.Fatalln(err.Error())
		}
		if err := conf.Init(); err != nil {
			log.Fatalln(err.Error())
		}
	}

	if opts.Watcher {
		go func() {
			for {
				err := runtimeViper.WatchRemoteConfig()
				if err != nil {
					log.Printf("unable to read remote config: %v", err)
					continue
				}

				for _, conf := range opts.Configs {
					if err := runtimeViper.Unmarshal(conf); err != nil {
						log.Fatalln(err.Error())
					}
					if err := conf.Init(); err != nil {
						log.Fatalln(err.Error())
					}
				}
			}
		}()
	}

	return EtcdConf{
		opts: opts.Configs,
	}
}
func (e EtcdConf) Init() error {
	return nil
}

func (e EtcdConf) Close() error {
	for _, o := range e.opts {
		_ = o.Close()
	}

	return nil
}
