package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	"github.com/ridwanakf/autobase-twitter/internal/usecase/archiveuc"
	"github.com/ridwanakf/autobase-twitter/internal/usecase/autobaseuc"
)

type Usecases struct {
	AutobaseUC internal.AutobaseUC
	ArchiveUC  internal.ArchiveUC
}

func newUsecases(repos *Repos, gateway *Gateways, cfg *config.Config) *Usecases {
	return &Usecases{
		AutobaseUC: autobaseuc.NewAutobaseUsecase(gateway.AutobaseGW),
		ArchiveUC:  archiveuc.NewArchiveUsecase(repos.ArchiveRepo),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
