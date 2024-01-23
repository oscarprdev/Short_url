package postgres

import (
	"database/sql"
	"errors"

	errorsC "short_url/pkg/features/shared/handlers"

	types "short_url/pkg/features/shared/types"
)

func ScanDbUser(row *sql.Row) (*types.DbUser, error) {
	var i types.DbUser
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
