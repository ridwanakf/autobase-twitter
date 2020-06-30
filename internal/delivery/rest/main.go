package rest

import (
	"github.com/labstack/echo"
	"github.com/ridwanakf/autobase-twitter/internal/app"
	"github.com/ridwanakf/autobase-twitter/internal/delivery/rest/server"
	"github.com/ridwanakf/autobase-twitter/internal/delivery/rest/service"
)

func initAPIHandler(eg *echo.Group, svc *service.Services) {
}

func Start(app *app.AutobaseApp) {
	srv := server.New()
	svc := service.GetServices(app)

	srv.GET("/", svc.IndexHandler)

	// API Handlers
	api := srv.Group("/api/v1")
	initAPIHandler(api, svc)

	server.Start(srv, &app.Cfg.Server)
}
