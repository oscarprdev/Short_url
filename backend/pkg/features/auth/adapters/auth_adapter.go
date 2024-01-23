package adapters

import (
	"time"

	api "short_url/pkg/api"

	"github.com/markbates/goth"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type DbUser struct {
	Id        string              `json:"id"`
	Email     openapi_types.Email `json:"email"`
	Picture   string              `json:"picture"`
	Name      string              `json:"name"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Token     string              `json:"token"`
	ExpiresAt time.Time           `json:"expires_at"`
}

func AdaptAuthInfoToDb(uInfo *goth.User) *DbUser {

	apiEmail := openapi_types.Email(uInfo.Email)
	currentTime := time.Now().UTC()

	return &DbUser{
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

func AdaptDbUserToApp(dbU *DbUser) *api.User {
	return &api.User{
		Id:        &dbU.Id,
		Email:     &dbU.Email,
		Picture:   &dbU.Picture,
		Name:      &dbU.Name,
		CreatedAt: &dbU.CreatedAt,
		UpdatedAt: &dbU.UpdatedAt,
	}
}
