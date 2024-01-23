package auth

import (
	"context"
	api "short_url/pkg/api"
)

type AuthRepo interface {
	AuthUser(ctx context.Context) (*[]api.User, error)
}
