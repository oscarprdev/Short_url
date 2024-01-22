package users

import (
	"database/sql"
	uinfra "short_url/pkg/features/users/infra"
	userUc "short_url/pkg/features/users/usecases"
	uusecases "short_url/pkg/features/users/usecases"
)

func WireUserUsecases(sql *sql.DB) *userUc.UserUsecases {
	repo := uinfra.ProvidePostgresRepository(sql)

	return uusecases.ProvideUserUsecases(repo)
}
