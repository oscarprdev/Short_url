package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	authAdapters "short_url/pkg/features/auth/adapters"

	errorsC "short_url/pkg/features/shared/handlers"
)

var insertPathQuery = "sql/queries/insert_user.sql"

func scanDbUser(row *sql.Row) (*authAdapters.DbUser, error) {
	var i authAdapters.DbUser
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Picture,
		&i.Name,
		&i.Token,
		&i.ExpiresAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, &errorsC.InternalError{
			Details: "Error retrieving user from DB",
		}
	}

	return &i, nil
}

func (pr *PostgresRepository) CreateNewUser(ctx context.Context, u *authAdapters.DbUser) error {
	fmt.Printf("CreateNewUser infra: %v\n", u)
	data, err := insertUser.ReadFile(insertPathQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Error on query",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(data),
		u.Id,
		u.CreatedAt,
		u.UpdatedAt,
		u.Email,
		u.Picture,
		u.Name,
		u.Token,
		u.ExpiresAt,
	)

	_, err = scanDbUser(row)
	if err != nil {
		return err
	}

	return nil
}

var selectPathQuery = "sql/queries/select_user_by_id.sql"

func (pr *PostgresRepository) GetUserById(ctx context.Context, id string) error {
	data, err := selectUserById.ReadFile(selectPathQuery)
	if err != nil {
		return &errorsC.InternalError{
			Details: "Select user query has an error",
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(data), id)

	_, err = scanDbUser(row)
	if err != nil {
		return err
	}

	return nil
}
