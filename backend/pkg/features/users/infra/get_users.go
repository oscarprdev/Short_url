package postgres

import (
	"context"

	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
)

var pathQuery = "sql/queries/elect_users.sql"

func (pr *PostgresRepository) GetUsers(ctx context.Context) (*[]api.User, error) {
	data, err := selectUsers.ReadFile(pathQuery)

	rows, err := pr.Db.QueryContext(ctx, string(data))
	if err != nil {
		return nil, &errors.BadRequestError{
			Details: "Error insterting data on database",
		}
	}

	defer rows.Close()

	var users []api.User
	for rows.Next() {
		var user api.User
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, &errors.BadRequestError{
				Details: "Error providing users from database",
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, &errors.BadRequestError{
			Details: "Some row on database is no",
		}
	}

	return &users, nil
}
