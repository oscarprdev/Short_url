package postgres

import (
	"context"
	"fmt"

	errors "short_url/pkg/features/shared/handlers"
	types "short_url/pkg/features/shared/types"
)

var pathQuery = "sql/queries/insert_url.sql"

func (pr *PostgresRepository) InsertUrl(ctx context.Context, rb *types.DbUrl) (*types.DbUrl, error) {
	query, err := insertUrlQuery.ReadFile(pathQuery)
	if err != nil {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading query: %v", err),
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), rb.Id, rb.CreatedAt, rb.UpdatedAt, rb.OriginalUrl, rb.ShortUrl, rb.Title, rb.ExpiresAt)

	var url types.DbUrl
	err = row.Scan(
		&url.Id,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.OriginalUrl,
		&url.ShortUrl,
		&url.Title,
		&url.ExpiresAt,
	)

	if err != nil {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading URL from database: %v", err),
		}
	}

	return &url, nil
}
