package postgres

import (
	"context"

	api "short_url/pkg/api"
)

func (pr *PostgresRepository) GetUsers(ctx context.Context) (*[]api.User, error) {
	rows, err := pr.Db.QueryContext(ctx, "SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []api.User
	for rows.Next() {
		var user api.User
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
