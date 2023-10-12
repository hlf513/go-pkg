package config

type Option func(*Options)

type Options struct {
	Name     string   `json:"name"`
	Endpoint string   `json:"endpoint"`
	Path     []string `json:"path"`
	Type     string   `json:"type"`
	Configs  []Config `json:"configs"`
	Watcher  bool     `json:"watcher"`
}

func NewOptions(opt ...Option) Options {
	opts := Options{
		Type: "json",
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}

func Path(path ...string) Option {
	return func(o *Options) {
		o.Path = append(o.Path, path...)
	}
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func Type(t string) Option {
	return func(o *Options) {
		o.Type = t
	}
}

func AddConfig(conf ...Config) Option {
	return func(o *Options) {
		o.Configs = conf
	}
}

func Endpoint(ep string) Option {
	return func(o *Options) {
		o.Endpoint = ep
	}
}

func Watcher(toggle bool) Option {
	return func(o *Options) {
		o.Watcher = toggle
	}
}
