package urls

import (
	"context"
	"fmt"
	api "short_url/pkg/api"
	adapters "short_url/pkg/features/urls/adapters"

	sharedAdapters "short_url/pkg/features/shared/adapters"
	errorsC "short_url/pkg/features/shared/handlers"
	validations "short_url/pkg/features/shared/validations"
)

func (uc *UrlUsecases) ShortUrlComplexUsecases(ctx context.Context, ou *adapters.OriginalUrl, url *api.Url, queryId, token string) error {
	dbu, err := uc.Repo.GetUserById(ctx, queryId)
	if err != nil {
		return &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	// Check if the user has permissions
	dbUser, err := validations.CheckUserIsValidated(ctx, dbu, queryId, token)
	if err != nil {
		return &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User is not authorized: %v", err),
		}
	}
	// SHORT URL
	su := shortUrl(ou)
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
	err = uc.RelateUrlWithUser(ctx, *ur.Id, dbUser.Id)
	if err != nil {
		return &errorsC.InternalError{
			Details: fmt.Sprintf("Error connecting url with user: %v", err),
		}
	}

	//Return url
	*url = *sharedAdapters.AdaptShortUrlToApp(suDB)

	return nil
}
