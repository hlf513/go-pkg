package openTelemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Tracer 获取注册的 Tracer
// Example:
// tracer := Tracer()
// ctx, span := tracer.Start(ctx,"uuid.service")
// ...
// span.End()
func Tracer(name ...string) trace.Tracer {
	var traceName = "github.com/hlf513/go-pkg/openTelemetry"
	if len(name) > 0 {
		traceName = name[0]
	}
	return otel.Tracer(traceName)
}

func Tag(span trace.Span, key, value string) {
	span.SetAttributes(attribute.String(key, value))
}

func Log(span trace.Span, message string) {
	span.AddEvent("log", trace.WithAttributes(attribute.String("message", message)))
}
