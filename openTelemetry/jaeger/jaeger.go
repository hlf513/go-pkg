package jaeger

import (
	"github.com/hlf513/go-pkg/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
	"os"
)

func TracerProvider(opt ...Option) *tracesdk.TracerProvider {
	opts := newOptions(opt...)
	// Create the Jaeger exporter
	exp, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(opts.AgentHost),
			jaeger.WithAgentPort(opts.AgentPort),
			jaeger.WithMaxPacketSize(opts.MaxPacketSize),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	var attrs = []attribute.KeyValue{
		semconv.ServiceNameKey.String(opts.ServiceName),
	}

	// Set the process field of jaeger
	if hostname, err := os.Hostname(); err == nil {
		attrs = append(attrs, attribute.String("hostname", hostname))
	}
	if ip, err := utils.HostIP(); err == nil {
		attrs = append(attrs, attribute.String("ip", ip.String()))
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(semconv.SchemaURL, attrs...)),
	)
	return tp
}
