package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	types "short_url/pkg/features/shared/types"
)

var updateTitleQuery = "sql/queries/update_title.sql"

func (pr *PostgresRepository) UpdateTitleUrl(ctx context.Context, url types.DbUrl) error {
	query, err := updateTitleUrlQuery.ReadFile(updateTitleQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Update url title query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), url.Title, url.Id)

	_, err = sql.ScanDbUrl(row)
	if err != nil {
		return err
	}

	return nil
}
