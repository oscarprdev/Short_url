package urls

import (
	"context"
	api "short_url/pkg/api"
	sharedAdapters "short_url/pkg/features/shared/adapters"

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

func (us *UrlUsecases) RelateUrlWithUser(ctx context.Context, urlId uuid.UUID, userId string) error {
	return us.Repo.RelateUrlWithUser(ctx, urlId, userId)
}
