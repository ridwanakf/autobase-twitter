package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
)

type Gateways struct {
	AutobaseGW internal.AutobaseGW
}

func newGateways(cfg *config.Config) *Gateways {
	return &Gateways{
	}
}

func (*Gateways) Close() []error {
	var errs []error
	return errs
}
