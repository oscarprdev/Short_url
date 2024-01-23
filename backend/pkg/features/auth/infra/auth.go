package postgres

import (
	"context"

	api "short_url/pkg/api"
)

var pathQuery = "sql/queries/elect_users.sql"

func (pr *PostgresRepository) AuthUser(ctx context.Context) (*[]api.User, error) {
	return nil, nil
}
