package urls

import (
	"context"
	"fmt"
	api "short_url/pkg/api"
	errorsC "short_url/pkg/features/shared/handlers"
	types "short_url/pkg/features/shared/types"
	validations "short_url/pkg/features/shared/validations"
	adapters "short_url/pkg/features/urls/adapters"
	"time"
)

func (uc *UrlUsecases) UrlTitleUsecases(ctx context.Context, urlBody *api.PostUrlTitleJSONBody, queryId string) (*api.Url, error) {
	dbu, err := uc.Repo.GetUserById(ctx, queryId)
	if err != nil {
		return nil, &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	// Check if the user has permissions
	_, err = validations.CheckUserIsValidated(ctx, dbu)
	if err != nil {
		return nil, &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User is not authorized: %v", err),
		}
	}

	url, err := uc.Repo.GetUrlById(ctx, *urlBody.UrlId)
	if err != nil {
		return nil, err
	}

	urlUpdated := types.DbUrl{
		Id:          url.Id,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   time.Now().UTC(),
		OriginalUrl: url.OriginalUrl,
		ShortUrl:    url.ShortUrl,
		Title:       *urlBody.UrlTitle,
		ExpiresAt:   url.ExpiresAt,
		Usage:       url.Usage,
	}

	err = uc.Repo.UpdateTitleUrl(ctx, urlUpdated)
	if err != nil {
		return nil, err
	}

	urlOK, err := adapters.AdaptUrldbToApp(&urlUpdated)
	if err != nil {
		return nil, err
	}

	return urlOK, nil
}
