package urls

import (
	"database/sql"
	uinfra "short_url/pkg/features/urls/infra"
	urlUc "short_url/pkg/features/urls/usecases"
)

func WireUrlsUsecases(sql *sql.DB) *urlUc.UrlUsecases {
	repo := uinfra.ProvidePostgresRepository(sql)

	return urlUc.ProvideUrlUsecases(repo)
}
