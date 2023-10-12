package mysql

const (
	Silent = "silent"
	Warn   = "warn"
	Error  = "error"
	Info   = "info"
)

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	opt := Options{
		Name:        "default",
		Host:        "127.0.0.1",
		Port:        3306,
		Username:    "user",
		Password:    "pwd",
		Database:    "dbname",
		MaxIdleConn: 10,
		MaxOpenConn: 100,
		MaxLifeTime: 1800,
		LogLevel:    Silent,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	Name        string `json:"name" mapstructure:"name"`
	Host        string `json:"host" mapstructure:"host"`
	Port        int32  `json:"port" mapstructure:"port"`
	Username    string `json:"username" mapstructure:"username"`
	Password    string `json:"password" mapstructure:"password"`
	Database    string `json:"database" mapstructure:"database"`
	MaxIdleConn int    `json:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn" mapstructure:"max_open_conn"`
	LogLevel    string `json:"log_level" mapstructure:"log_level"`
	MaxLifeTime int32  `json:"max_life_time" mapstructure:"max_life_time"`
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func LogLevel(level string) Option {
	return func(o *Options) {
		o.LogLevel = level
	}
}

func Host(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func Port(port int32) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func Username(username string) Option {
	return func(o *Options) {
		o.Username = username
	}
}

func Password(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}

func Database(database string) Option {
	return func(o *Options) {
		o.Database = database
	}
}

func MaxIdleConn(num int) Option {
	return func(o *Options) {
		if num > 0 {
			o.MaxIdleConn = num
		}
	}
}
func MaxOpenConn(num int) Option {
	return func(o *Options) {
		if num > 0 {
			o.MaxOpenConn = num
		}
	}
}

func MaxLifeTime(t int32) Option {
	return func(o *Options) {
		if t > 0 {
			o.MaxLifeTime = t
		}
	}
}
