package urls

import (
	"context"
	"fmt"
	api "short_url/pkg/api"
	sharedAdapters "short_url/pkg/features/shared/adapters"
	errorsC "short_url/pkg/features/shared/handlers"
	validations "short_url/pkg/features/shared/validations"

	types "short_url/pkg/features/shared/types"

	"github.com/google/uuid"
)

type UrlUsecases struct {
	Repo UrlRepo
}

func ProvideUrlUsecases(ur UrlRepo) *UrlUsecases {
	return &UrlUsecases{
		Repo: ur,
	}
}

func (us *UrlUsecases) InsertUrl(ctx context.Context, rb *types.DbUrl) (*api.Url, error) {
	dbUrl, err := us.Repo.InsertUrl(ctx, rb)
	if err != nil {
		return nil, err
	}

	return sharedAdapters.AdaptShortUrlToApp(dbUrl), nil
}

func (us *UrlUsecases) RemoveUrlById(ctx context.Context, id string, urlId string) error {
	dbu, err := us.Repo.GetUserById(ctx, id)
	if err != nil {
		return &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	// If user is not the default user
	if dbu.Id != "116176187754032784002" {
		// Check if the user has permissions
		_, err := validations.CheckUserIsValidated(ctx, dbu)
		if err != nil {
			return &errorsC.UnauthorizedError{
				Details: fmt.Sprintf("User is not authorized: %v", err),
			}
		}
	}

	err2 := us.Repo.RemoveUrlById(ctx, id, urlId)
	if err2 != nil {
		return err
	}

	return nil
}

func (us *UrlUsecases) RelateUrlWithUser(ctx context.Context, urlId uuid.UUID, userId string) error {
	return us.Repo.RelateUrlWithUser(ctx, urlId, userId)
}
