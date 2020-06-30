package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
)

type Usecases struct {
	AutobaseUC internal.AutobaseUC
}

func newUsecases(gateway *Gateways, cfg *config.Config) *Usecases {
	return &Usecases{
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
