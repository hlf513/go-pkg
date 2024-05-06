package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Request(ctx context.Context, opt ...Option) (responseBody []byte, httpCode int, err error) {
	opts := newOptions(opt...)

	var (
		request     *http.Request
		response    *http.Response
		requestBody io.Reader
		client      = http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
			Timeout:   opts.Timeout,
		}
	)

	// trace
	tr := otel.Tracer("github.com/hlf513/go-pkg/openTelemetry")
	ctx, span := tr.Start(ctx, "HTTP-"+strings.ToUpper(opts.Method))
	defer func() {
		if span != nil {
			span.SetAttributes(attribute.String("http.url", opts.Url))
			span.SetAttributes(attribute.Int("http.code", httpCode))
			if len(opts.Body) > 0 {
				span.AddEvent("log", trace.WithAttributes(attribute.String("http.request", string(opts.Body))))
			}
			span.AddEvent("log", trace.WithAttributes(attribute.String("http.response", string(responseBody))))
			span.End()
		}
	}()

	// request
	if len(opts.Body) > 0 {
		requestBody = bytes.NewReader(opts.Body)
	}
	if request, err = http.NewRequest(opts.Method, opts.Url, requestBody); err != nil {
		return
	}
	for k, v := range opts.Header {
		request.Header.Add(k, v)
	}
	ctx, cancelFunc := context.WithCancel(ctx) // 防止出现 goroutine 泄露
	defer cancelFunc()
	request = request.WithContext(ctx)
	if response, err = client.Do(request); err != nil {
		return
	}
	defer response.Body.Close()

	// response
	httpCode = response.StatusCode
	if responseBody, err = io.ReadAll(response.Body); err != nil {
		return
	}

	return
}
