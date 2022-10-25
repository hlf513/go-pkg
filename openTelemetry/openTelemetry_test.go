package openTelemetry

import (
	"context"
	"github.com/hlf513/go-pkg/openTelemetry/jaeger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	// init jaeger provider
	tp, err := jaeger.TracerProvider()
	assert.NoError(t, err)

	// register global and propagators
	tp = Register(tp)
	defer Shutdown(context.Background(), tp)
}
