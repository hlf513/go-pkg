package openTelemetry

import (
	"context"
	"errors"
	"github.com/hlf513/go-pkg/openTelemetry/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	// register global and propagators
	tp := Register(jaeger.TracerProvider())
	defer Shutdown(tp)

	// tracing
	tr := otel.Tracer("new trace")
	_, span := tr.Start(context.Background(), "api")
	defer span.End()
	// error
	span.SetStatus(codes.Error, "operationThatCouldFail failed")
	span.RecordError(errors.New("error message"))
	// tag
	span.SetAttributes(attribute.String("key", "value"))
	// log
	span.AddEvent("event_name", []trace.EventOption{
		trace.WithTimestamp(time.Now()),
		trace.WithStackTrace(true),
		trace.WithAttributes(attribute.String("sql", "select ... from ...")),
	}...)
}
