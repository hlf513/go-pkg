package redis

import "log"

var (
	DefaultName = "default"
	configs     = make(map[string]Options)
)

type Config struct {
	Redis []Options `json:"redis"`
}

func (c Config) Init() error {
	for _, r := range c.Redis {
		configs[r.Name] = r
		if err := Connect(
			Name(r.Name),
			Addr(r.Addr),
			Username(r.Username),
			Password(r.Password),
			DB(r.DB),
			MaxRetries(r.MaxRetries),
			DialTimeout(r.DialTimeout),
			ReadTimeout(r.ReadTimeout),
			WriteTimeout(r.WriteTimeout),
			PoolSize(r.PoolSize),
			PoolTimeout(r.PoolTimeout),
			Trace(r.Trace),
			Metrics(r.Metrics),
		); err != nil {
			log.Fatalf("[Redis] connect errors: %s; configure: %v", err.Error(), r)
			return err
		}

		log.Printf("[Redis] connected redis(%s)", r.Name)
	}

	return nil
}

func (c Config) Close() error {
	Close()
	return nil
}

func GetConfig(name ...string) Options {
	if len(name) == 0 {
		return configs[DefaultName]
	}
	return configs[name[0]]
}
