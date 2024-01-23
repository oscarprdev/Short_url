package auth

import (
	"context"
	"fmt"

	authAdapters "short_url/pkg/features/auth/adapters"

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
		dbUser := authAdapters.AdaptAuthInfoToDb(u)

		fmt.Printf("User mapped: %v\n", dbUser)

		if err := us.Repo.CreateNewUser(ctx, dbUser); err != nil {
			return err
		}

		return nil
	}

	fmt.Printf("Error here: %v\n", err)

	return nil
}
