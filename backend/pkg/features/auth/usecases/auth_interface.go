package auth

import (
	"context"
	adapters "short_url/pkg/features/auth/adapters"
)

type AuthRepo interface {
	CreateNewUser(ctx context.Context, u *adapters.DbUser) error
	GetUserById(ctx context.Context, id string) error
}
