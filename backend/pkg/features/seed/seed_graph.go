package seed

import (
	"database/sql"
)

func WireSeedUsecases(sql *sql.DB) *SeedUsecases {
	repo := ProvidePostgresRepository(sql)

	return ProvideSeedUsecase(repo)
}
