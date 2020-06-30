package service

import (
	"github.com/labstack/echo"
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app"
	"net/http"
)

type AutobaseService struct {
	uc internal.AutobaseUC
}

func NewAutobaseService(app *app.AutobaseApp) *AutobaseService {
	return &AutobaseService{
		uc: app.UseCases.AutobaseUC,
	}
}

func (s *AutobaseService) IndexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
