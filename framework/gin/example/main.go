package main

import (
	"github.com/hlf513/go-pkg/config"
	"github.com/hlf513/go-pkg/config/viper"
	"github.com/hlf513/go-pkg/database/go-redis"
	"github.com/hlf513/go-pkg/database/gorm/mysql"
	"github.com/hlf513/go-pkg/framework/gin"
	"github.com/hlf513/go-pkg/framework/gin/config"
	"github.com/hlf513/go-pkg/framework/gin/example/router"
	"github.com/hlf513/go-pkg/framework/gin/middleware/metrics"
	"github.com/hlf513/go-pkg/framework/gin/middleware/openTelemetry"
	"github.com/hlf513/go-pkg/framework/gin/middleware/pprof"
	"github.com/hlf513/go-pkg/log/zap"
	"log"

	"github.com/hlf513/go-pkg/openTelemetry/jaeger"
)

func main() {
	// new service
	s := gin.NewService(
		gin.Config(
			viper.NewFile(
				config.Name("example_config"),
				config.Path("./framework/gin/config"),
				config.AddConfig(
					new(redis.Config),
					new(zap.Config),
					new(jaeger.Config),
					new(mysql.Config),
					new(app.Config),
				),
			),
		),
		gin.Middleware(
			otelMdl.NewHandlerMiddleware(jaeger.GetConfig()),
			metrics.NewHandlerMiddleware(),
			pprof.NewHandlerMiddleware(),
		),
	)
	defer s.Stop()

	// init config and middleware
	if err := s.Init(); err != nil {
		log.Fatal(err.Error())
	}

	// register handler
	s.Router(router.RegisterHandler)

	// running
	if err := s.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
