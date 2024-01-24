package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	types "short_url/pkg/features/shared/types"
)

var selectPathQuery = "sql/queries/select_user_by_id.sql"

func (pr *PostgresRepository) GetUserById(ctx context.Context, id string) (*types.DbUser, error) {
	query, err := selectUserById.ReadFile(selectPathQuery)
	if err != nil {
		return nil, &errorsC.InternalError{
			Details: "Select user query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), id)

	user, err := sql.ScanDbUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}
