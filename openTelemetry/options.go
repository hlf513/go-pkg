package openTelemetry

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	opt := Options{
		Url: "http://localhost:14268/api/traces",
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	Name string `json:"name" mapstructure:"name"`
	Url  string `json:"url" mapstructure:"url"`
	Auth string `json:"auth" mapstructure:"auth"`
	Log  struct {
		Trace bool `json:"trace" mapstructure:"trace"`
		File  bool `json:"file" mapstructure:"file"`
	} `json:"log" mapstructure:"log"`
}

func ServiceName(serviceName string) Option {
	return func(o *Options) {
		o.Name = serviceName
	}
}

func TraceUrl(url string) Option {
	return func(o *Options) {
		o.Url = url
	}
}

func TraceLog(toggle bool) Option {
	return func(o *Options) {
		o.Log.Trace = toggle
	}
}

func FileLog(toggle bool) Option {
	return func(o *Options) {
		o.Log.File = toggle
	}
}

func Auth(auth string) Option {
	return func(o *Options) {
		o.Auth = auth
	}
}
