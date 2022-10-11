package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

var rds sync.Map

func Connect(opts ...Option) error {
	opt := newOptions(opts...)
	redisPool := &redis.Pool{
		MaxIdle:     opt.MaxIdleConn,
		MaxActive:   opt.MaxActiveConn,
		IdleTimeout: opt.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				opt.Server+":"+opt.Port,
				redis.DialConnectTimeout(opt.ConnectTimeout),
				redis.DialReadTimeout(opt.ReadTimeout),
				redis.DialWriteTimeout(opt.WriteTimeout),
			)
			if err != nil {
				return nil, err
			}
			if opt.Password != "" {
				if _, err := c.Do("AUTH", opt.Password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			if opt.SelectDB != 0 {
				if _, err := c.Do("SELECT", opt.SelectDB); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Wait: true,
	}

	rds.Store(opt.Name, redisPool)

	return nil
}

func Close() {
	rds.Range(func(key, value interface{}) bool {
		_ = value.(*redis.Pool).Close()
		return true
	})
}

func GetConn(name ...string) redis.Conn {
	var key = "default"
	if len(name) > 0 {
		key = name[0]
	}

	if db, ok := rds.Load(key); ok {
		return db.(*redis.Pool).Get()
	}

	log.Fatal("redis instance[" + key + "] not found")

	return nil
}
