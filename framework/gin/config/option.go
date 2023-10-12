package app

type Option func(*Options)

type Options struct {
	Name string `json:"name" mapstructure:"name"`
	Mode string `json:"mode" mapstructure:"mode"`
	Port string `json:"port" mapstructure:"port"`
	Env  string `json:"env" mapstructure:"env"`
}

func NewOptions(opts ...Option) Options {
	opt := Options{
		Name: "gin",
		Mode: "debug",
		Port: ":8888",
		Env:  "dev",
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}
func Mode(mode string) Option {
	return func(o *Options) {
		o.Mode = mode
	}
}
func Port(port string) Option {
	return func(o *Options) {
		o.Port = port
	}
}
func Env(env string) Option {
	return func(o *Options) {
		o.Env = env
	}
}
