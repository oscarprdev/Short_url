package seed

import (
	"context"
)

type SeedUsecases struct {
	Repo SeedRepository
}

func ProvideSeedUsecase(ur SeedRepository) *SeedUsecases {
	return &SeedUsecases{
		Repo: ur,
	}
}

func (uc *SeedUsecases) createSeed(ctx context.Context) error {
	err := uc.Repo.createUrlsTable(ctx)
	if err != nil {
		return err
	}

	err = uc.Repo.createUsersTable(ctx)
	if err != nil {
		return err
	}

	err = uc.Repo.createUrlUserTable(ctx)
	if err != nil {
		return err
	}

	return nil
}
