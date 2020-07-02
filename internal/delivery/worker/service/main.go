package service

import "github.com/ridwanakf/autobase-twitter/internal/app"

type Services struct {
	*WorkerService
}

func GetServices(app *app.AutobaseApp) *Services {
	return &Services{
		WorkerService: NewWorkerService(app),
	}
}
