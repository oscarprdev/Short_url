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

	"github.com/google/uuid"
)

func (uc *UrlUsecases) UrlUsageUsecases(ctx context.Context, urlId uuid.UUID, queryId string) (*api.Url, error) {
	dbu, err := uc.Repo.GetUserById(ctx, queryId)
	if err != nil {
		return nil, &errorsC.UnauthorizedError{
			Details: fmt.Sprintf("User id not valid: %v", err),
		}
	}

	// If user is not the default user
	if dbu.Id != "116176187754032784002" {
		// Check if the user has permissions
		_, err := validations.CheckUserIsValidated(ctx, dbu)
		if err != nil {
			return nil, &errorsC.UnauthorizedError{
				Details: fmt.Sprintf("User is not authorized: %v", err),
			}
		}
	}

	url, err := uc.Repo.GetUrlById(ctx, urlId)
	if err != nil {
		return nil, err
	}

	urlUpdated := types.DbUrl{
		Id:          url.Id,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   time.Now().UTC(),
		OriginalUrl: url.OriginalUrl,
		ShortUrl:    url.ShortUrl,
		Title:       url.Title,
		ExpiresAt:   url.ExpiresAt,
		Usage:       url.Usage + 1,
	}

	err = uc.Repo.UpdateUsageUrl(ctx, urlUpdated)
	if err != nil {
		return nil, err
	}

	urlOK, err := adapters.AdaptUrldbToApp(&urlUpdated)
	if err != nil {
		return nil, err
	}

	return urlOK, nil
}
