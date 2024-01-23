package adapters

import (
	api "short_url/pkg/api"
	"time"

	"github.com/google/uuid"
)


func AdaptShortUrlToDB(ou, su string) (*ShortUrlDBBody, error) {
	validOu, err := validateUrl(ou)
	if err != nil {
		return nil, err
	}

	title := "Default url title"
	now := time.Now().UTC()
	uuid := uuid.New()

	return &ShortUrlDBBody{
		Id:          uuid,
		OriginalUrl: *validOu.Ou,
		ShortUrl:    su,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       title,
		ExpiresAt:   now.Add(7 * 24 * time.Hour), // 1 week
	}, nil
}

func AdaptShortUrlToApp(subody *ShortUrlDBBody) *api.Url {
	return &api.Url{
		OriginalUrl: &subody.OriginalUrl,
		ShortUrl:    &subody.ShortUrl,
		UpdatedAt:   &subody.UpdatedAt,
		CreatedAt:   &subody.CreatedAt,
		ExpiresAt:   &subody.ExpiresAt,
		TitleUrl:    &subody.Title,
	}
}

func AdaptAuthUserInfoToDB(auo *AuthUserInfo) *AuthUserInfo {
	return &AuthUserInfo{
		Id:    auo.Id,
		Token: auo.Token,
	}
}
