package urls

import (
	"context"
	api "short_url/pkg/api"

	adapters "short_url/pkg/features/urls/adapters"
)

type UrlUsecases struct {
	Repo UrlRepo
}

func ProvideUrlUsecases(ur UrlRepo) *UrlUsecases {
	return &UrlUsecases{
		Repo: ur,
	}
}

func (us *UrlUsecases) InsertUrl(ctx context.Context, rb *adapters.ShortUrlDBBody) (*api.Url, error) {
	return us.Repo.InsertUrl(ctx, rb)
}
