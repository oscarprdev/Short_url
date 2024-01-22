package postgres

import (
	"context"
	"fmt"

	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
	adapters "short_url/pkg/features/urls/adapters"
)

var pathQuery = "sql/queries/insert_url_query.sql"

func (pr *PostgresRepository) InsertUrl(ctx context.Context, rb *adapters.ShortUrlDBBody) (*api.Url, error) {
	query, err := insertUrlQuery.ReadFile(pathQuery)
	if err != nil {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading query: %v", err),
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), rb.Id, rb.OriginalUrl, rb.ShortUrl)

	var url api.Url
	err = row.Scan(
		&url.Id,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.OriginalUrl,
		&url.ShortUrl,
		&url.TitleUrl,
		&url.ExpiresAt,
	)

	if err != nil {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading URL from database: %v", err),
		}
	}

	return &url, err
}
