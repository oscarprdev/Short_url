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
	dbUsers, err := uc.Repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var apiUsers []api.User
	for _, user := range *dbUsers {
		currentUser := user
		adaptedUser := adapters.AdaptDbUserToApp(&currentUser)
		apiUsers = append(apiUsers, *adaptedUser)
	}

	return &apiUsers, nil
}

func (uc *UserUsecases) DescribeUsers(ctx context.Context, id string) (*adapters.DescribeUser, error) {
	dbu, err := uc.Repo.GetUserById(ctx, id)
	if err != nil {
		return nil, &handlers.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	// Check if the user has permissions
	dbUser, err := validations.CheckUserIsValidated(ctx, dbu)
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
		currentUrl := url
		adaptedUrl := sharedAdapters.AdaptShortUrlToApp(&currentUrl)
		apiUrls = append(apiUrls, *adaptedUrl)
	}

	return &adapters.DescribeUser{
		User: apiUser,
		Urls: &apiUrls,
	}, nil
}
