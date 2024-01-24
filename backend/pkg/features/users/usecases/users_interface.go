package users

import (
	"context"
	api "short_url/pkg/api"
	"short_url/pkg/features/shared/types"
)

type UserRepo interface {
	GetUsers(ctx context.Context) (*[]api.User, error)
	GetUserById(ctx context.Context, id string) (*types.DbUser, error)
	JoinUserWithUrls(ctx context.Context, id string) (*[]types.DbUrl, error)
}
