package database

import "gorm.io/gorm/schema"

type QueryOption func(*QueryOptions)

type QueryOptions struct {
	Table           schema.Tabler
	BatchInsertSize int
	Fields          string
	Where           [2]string
	Group           string
	Order           string
	Join            string
	Having          string
	Offset          int
	Limit           int
}

func NewQueryOptions(opts ...QueryOption) QueryOptions {
	var opt = QueryOptions{
		BatchInsertSize: 1000,
		Fields:          "*",
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func Table(t schema.Tabler) QueryOption {
	return func(o *QueryOptions) {
		o.Table = t
	}
}
func BatchInsertSize(n int) QueryOption {
	return func(o *QueryOptions) {
		o.BatchInsertSize = n
	}
}
func Fields(s string) QueryOption {
	return func(o *QueryOptions) {
		o.Fields = s
	}
}
func Where(condition, value string) QueryOption {
	return func(o *QueryOptions) {
		o.Where = [2]string{condition, value}
	}
}
func Group(s string) QueryOption {
	return func(o *QueryOptions) {
		o.Group = s
	}
}
func Order(s string) QueryOption {
	return func(o *QueryOptions) {
		o.Order = s
	}
}
func Join(s string) QueryOption {
	return func(o *QueryOptions) {
		o.Join = s
	}
}
func Having(s string) QueryOption {
	return func(o *QueryOptions) {
		o.Having = s
	}
}
func Offset(n int) QueryOption {
	return func(o *QueryOptions) {
		o.Offset = n
	}
}
func Limit(n int) QueryOption {
	return func(o *QueryOptions) {
		o.Limit = n
	}
}