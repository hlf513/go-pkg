package openTelemetry

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"time"
)

func Register(tp *tracesdk.TracerProvider) *tracesdk.TracerProvider {
	// Register our TracerProvider as the global
	otel.SetTracerProvider(tp)
	// Propagators
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp
}

func Shutdown(ctx context.Context, tp *tracesdk.TracerProvider) {
	// Do not make the application hang when it is shutdown.
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
