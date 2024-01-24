package postgres

import (
	"context"
	errorsC "short_url/pkg/features/shared/handlers"
	types "short_url/pkg/features/shared/types"
)

var joinPathQuery = "sql/queries/join_user_url.sql"

func (pr *PostgresRepository) JoinUserWithUrls(ctx context.Context, id string) (*[]types.DbUrl, error) {
	query, err := joinUserUrl.ReadFile(joinPathQuery)
	if err != nil {
		return nil, &errorsC.InternalError{
			Details: "Select user query has an error",
		}
	}

	rows, err := pr.Db.QueryContext(ctx, string(query), id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []types.DbUrl

	for rows.Next() {
		var i types.DbUrl
		if err := rows.Scan(
			&i.Id,
			&i.OriginalUrl,
			&i.ShortUrl,
			&i.Title,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &items, nil
}
