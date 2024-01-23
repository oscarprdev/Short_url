package postgres

import (
	"context"
	"fmt"
	errorsC "short_url/pkg/features/shared/handlers"
	sql "short_url/pkg/features/shared/infra"
	types "short_url/pkg/features/shared/types"
)

var insertPathQuery = "sql/queries/insert_user.sql"

func (pr *PostgresRepository) CreateNewUser(ctx context.Context, u *types.DbUser) error {
	fmt.Printf("CreateNewUser infra: %v\n", u)
	query, err := insertUser.ReadFile(insertPathQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Error on query",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query),
		u.Id,
		u.CreatedAt,
		u.UpdatedAt,
		u.Email,
		u.Picture,
		u.Name,
		u.Token,
		u.ExpiresAt,
	)

	_, err = sql.ScanDbUser(row)
	if err != nil {
		return err
	}

	return nil
}
