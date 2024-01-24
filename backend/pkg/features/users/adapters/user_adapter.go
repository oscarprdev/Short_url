package adapters

import (
	api "short_url/pkg/api"
	"short_url/pkg/features/shared/types"
)

type DescribeUser struct {
	User *api.User  `json:"user"`
	Urls *[]api.Url `json:"urls"`
}

func AdaptDbUserToApp(dbUser *types.DbUser) *api.User {
	return &api.User{
		Id:        &dbUser.Id,
		Email:     &dbUser.Email,
		UpdatedAt: &dbUser.UpdatedAt,
		CreatedAt: &dbUser.CreatedAt,
		Name:      &dbUser.Name,
		Picture:   &dbUser.Picture,
	}
}
