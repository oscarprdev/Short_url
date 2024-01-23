package auth

import (
	"context"
	api "short_url/pkg/api"
)

type AuthUsecases struct {
	Repo AuthRepo
}

func ProvideAuthUsecases(ur AuthRepo) *AuthUsecases {
	return &AuthUsecases{
		Repo: ur,
	}
}

func (us *AuthUsecases) AuthUser(ctx context.Context) (*[]api.User, error) {
	return us.Repo.AuthUser(ctx)
}
