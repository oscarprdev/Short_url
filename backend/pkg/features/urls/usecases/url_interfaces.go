package urls

import (
	"context"
	api "short_url/pkg/api"
	adapters "short_url/pkg/features/urls/adapters"
)

type UrlRepo interface {
	InsertUrl(ctx context.Context, rb *adapters.ShortUrlDBBody) (*api.Url, error)
}
