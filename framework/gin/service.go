package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hlf513/go-pkg/framework/gin/config"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Service interface {
	Init() error
	Router(func(engine *gin.Engine))
	Start() error
	Stop() error
}

func NewService(opts ...Option) Service {
	s := new(service)
	s.opts = newOptions(opts...)
	return s
}

type service struct {
	opts   Options
	engine *gin.Engine
}

func (s *service) Init() error {
	// load config
	if err := s.opts.Config.Init(); err != nil {
		return err
	}
	// load middleware
	s.engine = gin.New()
	for _, m := range s.opts.Middleware {
		m(s.engine)
	}

	return nil
}

func (s *service) Router(fn func(engine *gin.Engine)) {
	fn(s.engine)
}

func (s *service) Start() error {
	srv := &http.Server{
		Addr:    app.GetConfig().Port,
		Handler: s.engine,
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Println("Gin Server running at http://0.0.0.0" + app.GetConfig().Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

	return nil
}

func (s *service) Stop() error {
	return s.opts.Config.Close()
}
