package auth

import (
	"database/sql"
	authI "short_url/pkg/features/auth/infra"
	authUc "short_url/pkg/features/auth/usecases"
)

func WireAuthUsecases(sql *sql.DB) *authUc.AuthUsecases {
	repo := authI.ProvidePostgresRepository(sql)

	return authUc.ProvideAuthUsecases(repo)
}
