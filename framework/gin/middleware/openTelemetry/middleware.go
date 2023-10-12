package otelMdl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hlf513/go-pkg/framework/gin/middleware"
	"github.com/hlf513/go-pkg/openTelemetry"
	"io"

	"github.com/hlf513/go-pkg/log/zap"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

func NewHandlerMiddleware(opts openTelemetry.Options) middleware.HandlerMiddleware {
	return func(engine *gin.Engine) {
		engine.Use(Middleware(opts))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Middleware 默认只记录请求地址等基本信息；根据配置文件可选记录请求日志
func Middleware(opts openTelemetry.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		tracerOpts := []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindServer),
		}
		spanName := c.FullPath()
		if spanName == "" {
			spanName = fmt.Sprintf("HTTP %s route not found", c.Request.Method)
		} else {
			rAttr := semconv.HTTPRouteKey.String(spanName)
			tracerOpts = append(tracerOpts, trace.WithAttributes(rAttr))
		}

		ctx, span := otel.Tracer("github.com/hlf513/go-pkg/framework/gin/middleware/openTelemetry").Start(ctx, spanName, tracerOpts...)
		defer span.End()
		c.Request = c.Request.WithContext(ctx)

		var (
			blw  *bodyLogWriter
			body []byte
			err  error
		)
		if opts.Log.Trace {
			if body, err = io.ReadAll(c.Request.Body); err != nil {
				span.RecordError(errors.New("http request body read exception"))
				span.SetStatus(codes.Error, err.Error())
			}
			span.AddEvent("log", trace.WithAttributes(attribute.String("request", string(body))))

			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			blw = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw
		}

		c.Next()

		if opts.Log.Trace {
			span.AddEvent("log", trace.WithAttributes(attribute.String("response", blw.body.String())))
		}

		if c.Writer.Status() != 200 {
			span.RecordError(errors.New("http code not 200"))
			span.SetStatus(codes.Error, "request exception")
		}
		span.SetAttributes(attribute.Int("http.code", c.Writer.Status()))

		if opts.Log.File {
			var logFmt = struct {
				ContextType string `json:"context_type"`
				Method      string
				Url         string
				Params      string
				Response    string
			}{
				ContextType: c.ContentType(),
				Method:      c.Request.Method,
				Url:         c.Request.URL.String(),
				Params:      string(body),
				Response:    blw.body.String(),
			}
			ginLog, _ := json.Marshal(logFmt)
			zap.Debug(c.Request.Context(), string(ginLog), "gin-log")
		}
	}
}
