package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	authAdapters "short_url/pkg/features/auth/adapters"

	api "short_url/pkg/api"
	errorsC "short_url/pkg/features/shared/handlers"
)

var insertPathQuery = "sql/queries/insert_user.sql"

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
	)

	var i api.User
	err = row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Picture,
		&i.Name,
	)

	return err
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

	var user authAdapters.DbUser
	err = row.Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Picture,
		&user.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}

		return &errorsC.InternalError{
			Details: "Error retrieving user from DB",
		}
	}

	return nil
}
