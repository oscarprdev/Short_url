package users

import (
	"context"
	"short_url/pkg/features/shared/types"
)

type UserRepo interface {
	GetUsers(ctx context.Context) (*[]types.DbUser, error)
	GetUserById(ctx context.Context, id string) (*types.DbUser, error)
	JoinUserWithUrls(ctx context.Context, id string) (*[]types.DbUrl, error)
}
