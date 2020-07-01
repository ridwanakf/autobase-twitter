package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	"github.com/ridwanakf/autobase-twitter/internal/usecase"
)

type Usecases struct {
	AutobaseUC internal.AutobaseUC
}

func newUsecases(gateway *Gateways, cfg *config.Config) *Usecases {
	return &Usecases{
		AutobaseUC: usecase.NewAutobaseUsecase(gateway.AutobaseGW),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
