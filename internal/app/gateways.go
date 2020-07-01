package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	"github.com/ridwanakf/autobase-twitter/internal/gateway"
)

type Gateways struct {
	AutobaseGW internal.AutobaseGW
}

func newGateways(cfg *config.Config) *Gateways {
	return &Gateways{
		AutobaseGW: gateway.NewAutobaseGateway(cfg.Twitter),
	}
}

func (*Gateways) Close() []error {
	var errs []error
	return errs
}
