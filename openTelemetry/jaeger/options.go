package jaeger

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		ServiceName:   "demo",
		AgentHost:     "localhost",
		AgentPort:     "6831",
		MaxPacketSize: 0,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Options struct {
	ServiceName   string
	AgentHost     string
	AgentPort     string
	MaxPacketSize int
}

func ServiceName(serviceName string) Option {
	return func(o *Options) {
		o.ServiceName = serviceName
	}
}

func AgentHost(host string) Option {
	return func(o *Options) {
		o.AgentHost = host
	}
}

func AgentPort(port string) Option {
	return func(o *Options) {
		o.AgentPort = port
	}
}

func MaxPacketSize(maxSize int) Option {
	return func(o *Options) {
		o.MaxPacketSize = maxSize
	}
}
