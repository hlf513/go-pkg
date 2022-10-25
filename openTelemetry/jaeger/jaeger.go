package jaeger

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func TracerProvider(opt ...Option) (*tracesdk.TracerProvider, error) {
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
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opts.ServiceName),
		)),
	)
	return tp, nil
}
