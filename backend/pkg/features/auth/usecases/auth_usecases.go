package auth

import (
	"context"

	authAdapters "short_url/pkg/features/auth/adapters"
	errorsC "short_url/pkg/features/shared/handlers"
	"short_url/pkg/features/shared/utils"

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
	user, err := us.Repo.GetUserById(ctx, u.UserID)
	if err != nil {
		// If type of error is from db then return error
		if internalErr, ok := err.(*errorsC.InternalError); ok {
			return internalErr
		}

		// If the user is not already registered on db then it is created
		dbUser := authAdapters.AdaptAuthInfoToDb(u)

		return us.Repo.CreateNewUser(ctx, dbUser)
	}

	// Update token if it is already expired
	err = utils.IsDateExpired(user.ExpiresAt)
	if err != nil {
		err = us.Repo.UpdateUser(ctx, u.ExpiresAt, u.AccessToken, user.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
