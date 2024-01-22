package users

import (
	"context"
	api "short_url/pkg/api"
)

type UserRepo interface {
	GetUsers(ctx context.Context) (*[]api.User, error)
}
