package jaeger

import (
	"github.com/hlf513/go-pkg/openTelemetry"
	"log"
)

var config openTelemetry.Options

type Config struct {
	Jaeger openTelemetry.Options `json:"jaeger"`
}

func (c Config) Init() error {
	config = c.Jaeger
	openTelemetry.Register(Jaeger(
		openTelemetry.ServiceName(c.Jaeger.Name),
		openTelemetry.TraceUrl(c.Jaeger.Url),
		openTelemetry.TraceLog(c.Jaeger.Log.Trace),
		openTelemetry.FileLog(c.Jaeger.Log.File),
	))
	log.Print("[Jaeger] init Jaeger configure")

	return nil
}

func (c Config) Close() error {
	if openTelemetry.IsRun() {
		openTelemetry.Shutdown(Jaeger())
	}
	return nil
}

func GetConfig() openTelemetry.Options {
	return config
}
