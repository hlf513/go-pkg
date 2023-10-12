package zap

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func otherFields(ctx context.Context, category string) []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.String("traceID", trace.SpanContextFromContext(ctx).TraceID().String()))
	return append(fields, zap.String("category", category))
}

func Info(ctx context.Context, msg, category string) {
	logger.Info(msg, otherFields(ctx, category)...)
}

func Warn(ctx context.Context, msg, category string) {
	logger.Warn(msg, otherFields(ctx, category)...)
}

func Error(ctx context.Context, msg, category string) {
	logger.Error(msg, otherFields(ctx, category)...)
}

func Debug(ctx context.Context, msg, category string) {
	logger.Debug(msg, otherFields(ctx, category)...)
}

func Fatal(ctx context.Context, msg, category string) {
	logger.Fatal(msg, otherFields(ctx, category)...)
}
