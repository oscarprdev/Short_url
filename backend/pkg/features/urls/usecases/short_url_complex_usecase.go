package urls

import (
	"context"
	"fmt"
	api "short_url/pkg/api"
	adapters "short_url/pkg/features/urls/adapters"

	sharedAdapters "short_url/pkg/features/shared/adapters"
	errorsC "short_url/pkg/features/shared/handlers"
	types "short_url/pkg/features/shared/types"
	validations "short_url/pkg/features/shared/validations"
)

func (uc *UrlUsecases) ShortUrlComplexUsecases(ctx context.Context, ou *adapters.OriginalUrl, url *api.Url, queryId, refererUrl string) error {
	dbu, err := uc.Repo.GetUserById(ctx, queryId)
	if err != nil {
		return &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	var validDbUser types.DbUser = *dbu

	// If user is not the default user
	if dbu.Id != "116176187754032784002" {
		// Check if the user has permissions
		dbUser, err := validations.CheckUserIsValidated(ctx, dbu)
		if err != nil {
			return &errorsC.UnauthorizedError{
				Details: fmt.Sprintf("User is not authorized: %v", err),
			}
		}

		validDbUser = *dbUser
	}

	// SHORT URL
	su := shortUrl(ou, refererUrl)
	suDB, err := adapters.AdaptShortUrlToDB(*ou.Ou, su)
	if err != nil {
		return &errorsC.BadRequestError{
			Details: fmt.Sprintf("Not valid fields: %v", err),
		}
	}

	// Insert short url to db
	ur, err := uc.InsertUrl(ctx, suDB)

	if err != nil {
		return &errorsC.InternalError{
			Details: fmt.Sprintf("Error creating new url on db: %v", err),
		}
	}

	// Insert short and user on db
	err = uc.RelateUrlWithUser(ctx, *ur.Id, validDbUser.Id)
	if err != nil {
		return &errorsC.InternalError{
			Details: fmt.Sprintf("Error connecting url with user: %v", err),
		}
	}

	//Return url
	*url = *sharedAdapters.AdaptShortUrlToApp(suDB)

	return nil
}
