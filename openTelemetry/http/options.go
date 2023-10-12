package http

import (
	"encoding/json"
	"net/url"
	"time"
)

const (
	ContentTypeJson = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

type Option func(*Options)

type Options struct {
	Url     string
	Method  string
	Header  map[string]string
	Timeout time.Duration
	Body    []byte
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Timeout: time.Second,
		Header:  make(map[string]string),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Url(u string, params ...map[string]string) Option {
	return func(o *Options) {
		if len(params) > 0 {
			p := url.Values{}
			up, _ := url.Parse(u)
			for k, v := range params[0] {
				p.Set(k, v)
			}
			up.RawQuery = p.Encode()
			u = up.String()
		}
		o.Url = u
	}
}

func Timeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func Method(method string) Option {
	return func(o *Options) {
		o.Method = method
	}
}

func Body(data []byte) Option {
	return func(o *Options) {
		o.Body = data
	}
}

func SetContentType(contentType string) Option {
	return func(o *Options) {
		o.Header["Content-Type"] = contentType
	}
}

func AddHeader(header map[string]string) Option {
	return func(o *Options) {
		for k, v := range header {
			o.Header[k] = v
		}
	}
}

func JsonBody(data any) Option {
	return func(o *Options) {
		body, _ := json.Marshal(data)
		o.Body = body
	}
}
