package adapters

import (
	api "short_url/pkg/api"
	"short_url/pkg/features/shared/types"
)

func AdaptShortUrlToApp(subody *types.DbUrl) *api.Url {
	return &api.Url{
		Id:          &subody.Id,
		OriginalUrl: &subody.OriginalUrl,
		ShortUrl:    &subody.ShortUrl,
		UpdatedAt:   &subody.UpdatedAt,
		CreatedAt:   &subody.CreatedAt,
		ExpiresAt:   &subody.ExpiresAt,
		TitleUrl:    &subody.Title,
	}
}
