package postgres

import (
	"database/sql"
	"errors"

	errorsC "short_url/pkg/features/shared/handlers"

	types "short_url/pkg/features/shared/types"
)

func ScanDbUrl(row *sql.Row) (*types.DbUrl, error) {
	var i types.DbUrl
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OriginalUrl,
		&i.ShortUrl,
		&i.Title,
		&i.ExpiresAt,
		&i.Usage,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, &errorsC.InternalError{
			Details: "Error retrieving url from DB",
		}
	}

	return &i, nil
}
