package service

import "github.com/ridwanakf/autobase-twitter/internal/app"

type Services struct {
	*AutobaseService
}

func GetServices(app *app.AutobaseApp) *Services {
	return &Services{
		AutobaseService: NewAutobaseService(app),
	}
}
