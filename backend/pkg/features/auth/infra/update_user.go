package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	"time"
)

var updatePathQuery = "sql/queries/update_user.sql"

func (pr *PostgresRepository) UpdateUser(ctx context.Context, expiresAt time.Time, token, id string) error {
	query, err := updateUser.ReadFile(updatePathQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Update user query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query),
		token,
		expiresAt,
		id,
	)

	_, err = sql.ScanDbUser(row)
	if err != nil {
		return err
	}

	return nil
}
