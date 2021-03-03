package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	"github.com/ridwanakf/autobase-twitter/internal/usecase/archiveuc"
	"github.com/ridwanakf/autobase-twitter/internal/usecase/autobaseuc"
	"github.com/ridwanakf/autobase-twitter/internal/usecase/rulesuc"
)

type Usecases struct {
	AutobaseUC internal.AutobaseUC
	ArchiveUC  internal.ArchiveUC
	RulesUC    internal.RulesUC
}

func newUsecases(repos *Repos, gateway *Gateways, cfg *config.Config) *Usecases {
	return &Usecases{
		AutobaseUC: autobaseuc.NewAutobaseUsecase(gateway.AutobaseGW),
		ArchiveUC:  archiveuc.NewArchiveUsecase(repos.ArchiveRepo),
		RulesUC:    rulesuc.NewRulesUsecase(gateway.AutobaseGW),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
