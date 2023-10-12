# Example
```go
func main() {
	openTelemetry.Register(jaeger.Jaeger(
		openTelemetry.ServiceName("go-redis"),
		openTelemetry.TraceUrl("http://localhost:14268/api/traces"),
	))
	defer openTelemetry.Shutdown(jaeger.Jaeger())

	if err := Connect(); err != nil {
		panic(err)
	}
	defer Close()

	ctx := context.Background()
	tracer := openTelemetry.Tracer()
	ctx, span := tracer.Start(ctx, "go-redis")
	defer span.End()

	cli := GetClient()
	cli.Get(ctx, "key")
}
```