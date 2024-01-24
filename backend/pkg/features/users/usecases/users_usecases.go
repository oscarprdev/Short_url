package users

import (
	"context"
	"fmt"
	api "short_url/pkg/api"

	sharedAdapters "short_url/pkg/features/shared/adapters"
	handlers "short_url/pkg/features/shared/handlers"
	"short_url/pkg/features/shared/validations"
	adapters "short_url/pkg/features/users/adapters"
)

type UserUsecases struct {
	Repo UserRepo
}

func ProvideUserUsecases(ur UserRepo) *UserUsecases {
	return &UserUsecases{
		Repo: ur,
	}
}

func (uc *UserUsecases) GetUsers(ctx context.Context) (*[]api.User, error) {
	return uc.Repo.GetUsers(ctx)
}

func (uc *UserUsecases) DescribeUsers(ctx context.Context, token, id string) (*adapters.DescribeUser, error) {
	dbu, err := uc.Repo.GetUserById(ctx, id)
	if err != nil {
		return nil, &handlers.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	// Check if the user has permissions
	dbUser, err := validations.CheckUserIsValidated(ctx, dbu, id, token)
	if err != nil {
		return nil, &handlers.UnauthorizedError{
			Details: fmt.Sprintf("User is not authorized: %v", err),
		}
	}

	urls, err := uc.Repo.JoinUserWithUrls(ctx, dbUser.Id)
	if err != nil {
		return nil, &handlers.InternalError{
			Details: fmt.Sprintf("Internal error: %v", err),
		}
	}

	apiUser := adapters.AdaptDbUserToApp(dbUser)

	var apiUrls []api.Url
	for _, url := range *urls {
		apiUrls = append(apiUrls, *sharedAdapters.AdaptShortUrlToApp(&url))
	}

	return &adapters.DescribeUser{
		User: apiUser,
		Urls: &apiUrls,
	}, nil
}
