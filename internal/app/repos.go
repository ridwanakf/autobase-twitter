package app

import (
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	firebase2 "github.com/ridwanakf/autobase-twitter/internal/repo/db/firebase"
)

type Repos struct {
	ArchiveRepo internal.ArchiveRepo
}

func newRepos(cfg *config.Config) (*Repos, error) {
	// Check if archive feature is enabled
	if !cfg.UseArchive {
		return &Repos{
			ArchiveRepo: nil,
		}, nil
	}

	// TODO: Change this initialization according your own choice of database
	repo, err := firebase2.NewRealtimeDatabase(cfg.Databases.Firebase)
	if err != nil {
		return nil, err
	}

	return &Repos{
		ArchiveRepo: repo,
	}, nil
}

func (*Repos) Close() []error {
	var errs []error
	return errs
}
