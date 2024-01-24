package adapters

import (
	api "short_url/pkg/api"
	types "short_url/pkg/features/shared/types"
	"strings"
	"time"

	"github.com/google/uuid"
)

func AdaptShortUrlToDB(ou, su string) (*types.DbUrl, error) {
	validOu, err := validateUrl(ou)
	if err != nil {
		return nil, err
	}

	urlSplitted := strings.Split(su, "/")
	title := urlSplitted[len(urlSplitted)-1]

	now := time.Now().UTC()
	uuid := uuid.New()

	return &types.DbUrl{
		Id:          uuid,
		OriginalUrl: *validOu.Ou,
		ShortUrl:    su,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       title,
		Usage:       0,
		ExpiresAt:   now.Add(7 * 24 * time.Hour), // 1 week
	}, nil
}

func AdaptAuthUserInfoToDB(auo *AuthUserInfo) *AuthUserInfo {
	return &AuthUserInfo{
		Id:    auo.Id,
		Token: auo.Token,
	}
}

func AdaptUrldbToApp(urldb *types.DbUrl) (*api.Url, error) {
	return &api.Url{
		Id:          &urldb.Id,
		OriginalUrl: &urldb.OriginalUrl,
		ShortUrl:    &urldb.ShortUrl,
		CreatedAt:   &urldb.CreatedAt,
		UpdatedAt:   &urldb.UpdatedAt,
		TitleUrl:    &urldb.Title,
		Usage:       &urldb.Usage,
		ExpiresAt:   &urldb.ExpiresAt,
	}, nil
}
