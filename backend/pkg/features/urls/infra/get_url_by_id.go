package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	types "short_url/pkg/features/shared/types"

	"github.com/google/uuid"
)

var selectUrlPathQuery = "sql/queries/select_url_by_id.sql"

func (pr *PostgresRepository) GetUrlById(ctx context.Context, id uuid.UUID) (*types.DbUrl, error) {
	query, err := selectUrlById.ReadFile(selectUrlPathQuery)
	if err != nil {
		return nil, &errorsC.InternalError{
			Details: "Select url query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), id)

	url, err := sql.ScanDbUrl(row)
	if err != nil {
		return nil, err
	}

	return url, nil
}
