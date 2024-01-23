package auth

import (
	"context"

	authAdapters "short_url/pkg/features/auth/adapters"
	errorsC "short_url/pkg/features/shared/handlers"

	"github.com/markbates/goth"
)

type AuthUsecases struct {
	Repo AuthRepo
}

func ProvideAuthUsecases(ur AuthRepo) *AuthUsecases {
	return &AuthUsecases{
		Repo: ur,
	}
}

func (us *AuthUsecases) AuthUser(ctx context.Context, u *goth.User) error {
	err := us.Repo.GetUserById(ctx, u.UserID)

	if err != nil {
		if internalErr, ok := err.(*errorsC.InternalError); ok {
			return internalErr
		}

		dbUser := authAdapters.AdaptAuthInfoToDb(u)

		return us.Repo.CreateNewUser(ctx, dbUser)
	}

	return nil
}
