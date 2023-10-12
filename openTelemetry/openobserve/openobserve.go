package openObserve

import (
	"context"
	"fmt"
	"github.com/hlf513/go-pkg/openTelemetry"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
	"net/url"
	"strings"
	"sync"
)

var (
	openObserveProvider *tracesdk.TracerProvider
	openObserveOnce     sync.Once
)

func OpenObserve(opt ...openTelemetry.Option) *tracesdk.TracerProvider {
	openObserveOnce.Do(func() {
		opts := openTelemetry.NewOptions(opt...)

		u, err := url.Parse(opts.Url)
		if err != nil {
			log.Fatalf("openObserve url error:%s", err.Error())
		}

		var otlpOptions []otlptracehttp.Option
		otlpOptions = append(otlpOptions,
			otlptracehttp.WithEndpoint(u.Host),
			otlptracehttp.WithURLPath(u.Path),
			otlptracehttp.WithHeaders(map[string]string{
				"Authorization": fmt.Sprintf("Basic %s", opts.Auth),
			}),
		)
		if strings.ToLower(u.Scheme) == "http" {
			otlpOptions = append(otlpOptions, otlptracehttp.WithInsecure())
		}
		otlpHTTPExporter, err := otlptracehttp.New(context.TODO(), otlpOptions...)
		if err != nil {
			log.Println("openObserve error:", err.Error())
			return
		}

		res := resource.NewWithAttributes(
			semconv.SchemaURL,
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(opts.Name),
			//semconv.ServiceVersionKey.String("0.0.1"),
			//attribute.String("environment", "test"),
		)

		openObserveProvider = tracesdk.NewTracerProvider(
			tracesdk.WithResource(res),
			tracesdk.WithBatcher(otlpHTTPExporter),
			// Sampler
			//tracesdk.WithSampler(tracesdk.AlwaysSample()),
			//tracesdk.WithSampler(tracesdk.NeverSample()),
			//tracesdk.WithSampler(tracesdk.TraceIDRatioBased(0.5)),
			//tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.AlwaysSample())), // default
		)
	})

	return openObserveProvider
}
