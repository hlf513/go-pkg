package config

type Config interface {
	Init() error
	Close() error
}
