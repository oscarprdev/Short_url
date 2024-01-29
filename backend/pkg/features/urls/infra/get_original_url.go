package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	types "short_url/pkg/features/shared/types"
)

var selectOriginalPathQuery = "sql/queries/select_original_url.sql"

func (pr *PostgresRepository) GetOriginalUrl(ctx context.Context, shortUrl string) (*types.DbUrl, error) {
	query, err := selectOriginalQuery.ReadFile(selectOriginalPathQuery)
	if err != nil {
		return nil, &errorsC.InternalError{
			Details: "Select original url query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), shortUrl)

	url, err := sql.ScanDbUrl(row)
	if err != nil {
		return nil, err
	}

	return url, nil
}
