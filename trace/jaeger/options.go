package jaeger

type Option func(*Options)

type Options struct {
	ServiceName   string
	Address       string
	MaxPacketSize int
}
