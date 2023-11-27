package mysql

import (
	"log"
)

var (
	Default = "default"
	configs = make(map[string]Options)
)

type Config struct {
	MySQL []Options `json:"mysql"`
}

func (c Config) Init() error {
	for _, m := range c.MySQL {
		configs[m.Name] = m
		if err := Connect(
			Name(m.Name),
			Host(m.Host),
			Port(m.Port),
			Username(m.Username),
			Password(m.Password),
			Database(m.Database),
			LogLevel(m.LogLevel),
			MaxIdleConn(m.MaxIdleConn),
			MaxOpenConn(m.MaxOpenConn),
			MaxLifeTime(m.MaxLifeTime),
			Metrics(m.Metrics),
			Trace(m.Trace),
		); err != nil {
			log.Fatalf("[MySQL] connect errors: %s; configure: %v", err.Error(), m)
			return err
		}
		if m.LogLevel == "debug" {
			log.Printf("[MySQL] connected MySQL(%s),configure: %#v", m.Name, m)
		} else {
			log.Printf("[MySQL] connected MySQL(%s)", m.Name)
		}
	}

	return nil
}

func (c Config) Close() error {
	return nil
}

func GetConfig(name ...string) Options {
	if len(name) == 0 {
		return configs[Default]
	}
	return configs[name[0]]
}
