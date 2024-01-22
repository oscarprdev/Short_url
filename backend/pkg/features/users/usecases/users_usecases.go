package users

import (
	"context"
	api "short_url/pkg/api"
)

type UserUsecases struct {
	Repo UserRepo
}

func ProvideUserUsecases(ur UserRepo) *UserUsecases {
	return &UserUsecases{
		Repo: ur,
	}
}

func (us *UserUsecases) GetUsers(ctx context.Context) (*[]api.User, error) {
	return us.Repo.GetUsers(ctx)
}
