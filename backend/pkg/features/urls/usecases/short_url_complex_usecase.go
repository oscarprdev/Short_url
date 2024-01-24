package urls

import (
	"context"
	"fmt"
	api "short_url/pkg/api"
	adapters "short_url/pkg/features/urls/adapters"

	errorsC "short_url/pkg/features/shared/handlers"
	types "short_url/pkg/features/shared/types"
	utils "short_url/pkg/features/shared/utils"
)

func (us *UrlUsecases) checkUserIsValidated(ctx context.Context, queryId, token string) (*types.DbUser, error) {
	dbUser, err := us.Repo.GetUserById(ctx, queryId)
	if err != nil {
		return nil, &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	for i, char := range token {
		if string(dbUser.Token[i]) != string(byte(char)) {
			return nil, &errorsC.UnauthorizedError{
				Details: fmt.Sprintf("User token not valid: %v", err),
			}
		}
	}

	err = utils.IsDateExpired(dbUser.ExpiresAt)
	if err != nil {
		return nil, &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("Token expired: %v", err),
		}
	}

	return dbUser, nil
}

func (uc *UrlUsecases) ShortUrlComplexUsecases(ctx context.Context, ou *adapters.OriginalUrl, url *api.Url, queryId, token string) error {
	// Check if the user has permissions
	dbUser, err := uc.checkUserIsValidated(ctx, queryId, token)
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
	*url = *adapters.AdaptShortUrlToApp(suDB)

	return nil
}
