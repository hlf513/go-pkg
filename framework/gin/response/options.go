package response

type Option func(*Options)

type Options struct {
	HttpCode    int
	SuccessCode int32
	ErrorCode   int32

	CodeName    string
	MessageName string
	DataName    string

	SuccessMessage string
}

func newOptions(opts ...Option) Options {
	opt := Options{
		HttpCode:       200,
		SuccessCode:    0,
		ErrorCode:      1,
		CodeName:       "code",
		MessageName:    "msg",
		SuccessMessage: "success",
		DataName:       "data",
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func SuccessCode(code int32) Option {
	return func(o *Options) {
		o.SuccessCode = code
	}
}
func ErrorCode(code int32) Option {
	return func(o *Options) {
		o.ErrorCode = code
	}
}
func HttpCode(code int) Option {
	return func(o *Options) {
		o.HttpCode = code
	}
}

func MessageName(name string) Option {
	return func(o *Options) {
		o.MessageName = name
	}
}
func SuccessMessage(message string) Option {
	return func(o *Options) {
		o.SuccessMessage = message
	}
}
func DataName(name string) Option {
	return func(o *Options) {
		o.DataName = name
	}
}
func CodeName(name string) Option {
	return func(o *Options) {
		o.CodeName = name
	}
}
