package urls

import (
	"context"
	api "short_url/pkg/api"
	adapters "short_url/pkg/features/urls/adapters"

	types "short_url/pkg/features/shared/types"

	"github.com/google/uuid"
)

type UrlRepo interface {
	InsertUrl(ctx context.Context, rb *adapters.ShortUrlDBBody) (*api.Url, error)
	GetUserById(ctx context.Context, id string) (*types.DbUser, error)
	RelateUrlWithUser(ctx context.Context, urlId uuid.UUID, userId string) error
}
