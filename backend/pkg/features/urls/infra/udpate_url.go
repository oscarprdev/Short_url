package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	types "short_url/pkg/features/shared/types"
)

var updateQuery = "sql/queries/update_url.sql"

func (pr *PostgresRepository) UpdateUrl(ctx context.Context, url types.DbUrl) error {
	query, err := updateUrlQuery.ReadFile(updateQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Update url query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), url.Usage, url.UpdatedAt, url.Id)

	_, err = sql.ScanDbUrl(row)
	if err != nil {
		return err
	}

	return nil
}
