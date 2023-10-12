package openObserve

import (
	"github.com/hlf513/go-pkg/openTelemetry"
	"log"
)

var config openTelemetry.Options

type Config struct {
	OpenObserve openTelemetry.Options `json:"openObserve"`
}

func (c Config) Init() error {
	config = c.OpenObserve
	openTelemetry.Register(OpenObserve(
		openTelemetry.Auth(c.OpenObserve.Auth),    // https://openobserve.ai/docs/ingestion/traces/#credentials
		openTelemetry.TraceUrl(c.OpenObserve.Url), // https://openobserve.ai/docs/ingestion/traces/#self-hosted-openobserve
		openTelemetry.ServiceName(c.OpenObserve.Name),
		openTelemetry.FileLog(c.OpenObserve.Log.File),
		openTelemetry.TraceLog(c.OpenObserve.Log.Trace),
	))
	log.Print("[openObserve] init configure")

	return nil
}

func (c Config) Close() error {
	if openTelemetry.IsRun() {
		openTelemetry.Shutdown(OpenObserve())
	}
	return nil
}

func GetConfig() openTelemetry.Options {
	return config
}
