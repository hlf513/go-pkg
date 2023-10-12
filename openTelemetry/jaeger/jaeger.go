package jaeger

import (
	"github.com/hlf513/go-pkg/openTelemetry"
	"github.com/uber/jaeger-client-go/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
	"os"
	"sync"
)

var jaegerProvider *tracesdk.TracerProvider
var jaegerOnce sync.Once

func Jaeger(opt ...openTelemetry.Option) *tracesdk.TracerProvider {
	jaegerOnce.Do(func() {
		opts := openTelemetry.NewOptions(opt...)
		// Create the Jaeger exporter
		exp, err := jaeger.New(
			jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(opts.Url)),
		)
		if err != nil {
			log.Println("jaeger error: ", err.Error())
			return
		}
		var attrs = []attribute.KeyValue{
			semconv.ServiceNameKey.String(opts.Name),
		}

		// Set the process field of jaeger
		if hostname, err := os.Hostname(); err == nil {
			attrs = append(attrs, attribute.String("hostname", hostname))
		}
		if ip, err := utils.HostIP(); err == nil {
			attrs = append(attrs, attribute.String("ip", ip.String()))
		}

		jaegerProvider = tracesdk.NewTracerProvider(
			// Always be sure to batch in production.
			tracesdk.WithBatcher(exp),
			// Record information about this application in a Resource.
			tracesdk.WithResource(resource.NewWithAttributes(semconv.SchemaURL, attrs...)),
			// Sampler
			//tracesdk.WithSampler(tracesdk.AlwaysSample()),
			//tracesdk.WithSampler(tracesdk.NeverSample()),
			//tracesdk.WithSampler(tracesdk.TraceIDRatioBased(0.5)),
			//tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.AlwaysSample())), // default
		)
	})

	return jaegerProvider
}
