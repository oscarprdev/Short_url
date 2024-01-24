package urls

import (
	"context"

	types "short_url/pkg/features/shared/types"

	"github.com/google/uuid"
)

type UrlRepo interface {
	InsertUrl(ctx context.Context, rb *types.DbUrl) (*types.DbUrl, error)
	GetUserById(ctx context.Context, id string) (*types.DbUser, error)
	RelateUrlWithUser(ctx context.Context, urlId uuid.UUID, userId string) error
	GetUrlById(ctx context.Context, id uuid.UUID) (*types.DbUrl, error)
	UpdateUsageUrl(ctx context.Context, url types.DbUrl) error
	UpdateTitleUrl(ctx context.Context, url types.DbUrl) error
}
