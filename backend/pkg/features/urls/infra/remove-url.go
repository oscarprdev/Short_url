package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
)

var removeQuery = "sql/queries/remove_url.sql"

func (pr *PostgresRepository) RemoveUrlById(ctx context.Context, id string, urlId string) error {
	query, err := removeUrlQuery.ReadFile(removeQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Remove url query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), urlId)

	_, err = sql.ScanDbUrl(row)
	if err != nil {
		return err
	}

	return nil
}
