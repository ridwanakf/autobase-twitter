package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	md "github.com/ridwanakf/autobase-twitter/internal/delivery/rest/middleware"
)

// New instantates new Echo server
func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger(),
		middleware.Recover(),
		middleware.Secure(),
		md.CORS(),
		md.Headers())

	e.Validator = &CustomValidator{V: validator.New()}
	custErr := &customErrHandler{e: e}
	e.HTTPErrorHandler = custErr.handler
	e.Binder = &CustomBinder{b: &echo.DefaultBinder{}}

	return e
}

// Start starts echo server
func Start(e *echo.Echo, cfg *config.Server) {
	s := &http.Server{
		Addr:         ":" + cfg.Port,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}
	e.Debug = cfg.Debug

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
