package adapters

import (
	"time"

	api "short_url/pkg/api"

	"github.com/markbates/goth"

	openapi_types "github.com/oapi-codegen/runtime/types"

	types "short_url/pkg/features/shared/types"
)

func AdaptAuthInfoToDb(uInfo *goth.User) *types.DbUser {

	apiEmail := openapi_types.Email(uInfo.Email)
	currentTime := time.Now().UTC()

	return &types.DbUser{
		Id:        uInfo.UserID,
		Email:     apiEmail,
		Picture:   uInfo.AvatarURL,
		Name:      uInfo.Name,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Token:     uInfo.AccessToken,
		ExpiresAt: uInfo.ExpiresAt,
	}
}

func AdaptDbUserToApp(dbU *types.DbUser) *api.User {
	return &api.User{
		Id:        &dbU.Id,
		Email:     &dbU.Email,
		Picture:   &dbU.Picture,
		Name:      &dbU.Name,
		CreatedAt: &dbU.CreatedAt,
		UpdatedAt: &dbU.UpdatedAt,
	}
}
