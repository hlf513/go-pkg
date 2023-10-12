package query

import "gorm.io/gorm/schema"

type Option func(*Options)

type Options struct {
	Table           schema.Tabler
	BatchInsertSize int
	Select          string
	Where           [][]any
	Group           string
	Order           string
	Join            string
	Having          string
	Offset          int
	Limit           int
	HardDelete      bool
}

func newOptions(opts ...Option) Options {
	var opt = Options{
		BatchInsertSize: 1000,
		Select:          "*",
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func HardDelete() Option {
	return func(o *Options) {
		o.HardDelete = true
	}
}

func Table(t schema.Tabler) Option {
	return func(o *Options) {
		o.Table = t
	}
}
func BatchInsertSize(n int) Option {
	return func(o *Options) {
		o.BatchInsertSize = n
	}
}

func Select(s string) Option {
	return func(o *Options) {
		o.Select = s
	}
}
func Where(condition string, value ...any) Option {
	return func(o *Options) {
		where := []any{condition}
		where = append(where, value...)
		o.Where = append(o.Where, where)
	}
}
func Group(s string) Option {
	return func(o *Options) {
		o.Group = s
	}
}
func Order(s string) Option {
	return func(o *Options) {
		o.Order = s
	}
}
func Join(s string) Option {
	return func(o *Options) {
		o.Join = s
	}
}
func Having(s string) Option {
	return func(o *Options) {
		o.Having = s
	}
}
func Offset(n int) Option {
	return func(o *Options) {
		o.Offset = n
	}
}
func Limit(n int) Option {
	return func(o *Options) {
		o.Limit = n
	}
}
