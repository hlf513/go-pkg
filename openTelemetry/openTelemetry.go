package openTelemetry

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"
)

func Register(tp *tracesdk.TracerProvider) *tracesdk.TracerProvider {
	if tp == nil {
		return nil
	}
	// Register our TracerProvider as the global
	otel.SetTracerProvider(tp)
	// Register Propagators
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp
}

func Shutdown(tp *tracesdk.TracerProvider) {
	if tp == nil {
		return
	}
	// Do not make the application hang when it is shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func IsRun() bool {
	return otel.GetTracerProvider() != nil
}

func GetTraceIdFromContext(ctx context.Context) string {
	return trace.SpanContextFromContext(ctx).TraceID().String()
}
