package zap

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

type Option func(*Options)

type Options struct {
	LogPath        string `json:"log_path" mapstructure:"log_path"`
	FileMaxSize    int    `json:"file_max_size" mapstructure:"file_max_size" `
	FileMaxBackups int    `json:"file_max_backups" mapstructure:"file_max_backups"`
	FileMaxAge     int    `json:"file_max_age" mapstructure:"file_max_age"`
	FileCompress   bool   `json:"file_compress" mapstructure:"file_compress"`
	LogLevel       string `json:"log_level" mapstructure:"log_level"`
}

func newOption(opts ...Option) Options {
	opt := Options{
		FileMaxSize:    10,
		FileMaxBackups: 10,
		FileMaxAge:     7,
		FileCompress:   false,
		LogLevel:       LevelDebug,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func LogPath(path string) Option {
	return func(o *Options) {
		o.LogPath = path
	}
}
func FileMaxSize(size int) Option {
	return func(o *Options) {
		o.FileMaxSize = size
	}
}
func FileMaxBackups(num int) Option {
	return func(o *Options) {
		o.FileMaxBackups = num
	}
}
func FileMaxAge(days int) Option {
	return func(o *Options) {
		o.FileMaxAge = days
	}
}
func FileCompress(compress bool) Option {
	return func(o *Options) {
		o.FileCompress = compress
	}
}
func LogLevel(level string) Option {
	return func(o *Options) {
		o.LogLevel = level
	}
}
