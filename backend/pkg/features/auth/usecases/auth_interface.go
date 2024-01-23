package auth

import (
	"context"
	types "short_url/pkg/features/shared/types"
	"time"
)

type AuthRepo interface {
	CreateNewUser(ctx context.Context, u *types.DbUser) error
	GetUserById(ctx context.Context, id string) (*types.DbUser, error)
	UpdateUser(ctx context.Context, expiresAt time.Time, token, id string) error
}
