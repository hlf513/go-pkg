package redis

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	goRedis "github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var rds sync.Map

func Connect(opts ...Option) error {
	opt := newOptions(opts...)
	rdb := goRedis.NewClient(&goRedis.Options{
		Addr:         opt.Addr,
		Username:     opt.Username,
		Password:     opt.Password,
		DB:           opt.DB,
		MaxRetries:   opt.MaxRetries,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		PoolTimeout:  opt.PoolTimeout,
	})
	if opt.Trace {
		if err := redisotel.InstrumentTracing(rdb); err != nil {
			return err
		}
	}
	if opt.Metrics {
		if err := redisotel.InstrumentMetrics(rdb); err != nil {
			return err
		}
	}

	rds.Store(opt.Name, rdb)

	return nil
}

func Close() {
	rds.Range(func(key, value any) bool {
		if value != nil {
			_ = value.(*goRedis.Client).Close()
		}
		return true
	})
}

func GetClient(name ...string) *goRedis.Client {
	var key = DefaultName
	if len(name) > 0 {
		key = name[0]
	}

	if db, ok := rds.Load(key); ok {
		return db.(*goRedis.Client)
	}

	log.Fatal("[Fatal] redis client[" + key + "] not found")

	return nil
}
