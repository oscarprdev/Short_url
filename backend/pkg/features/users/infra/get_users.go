package postgres

import (
	"context"
	"fmt"

	errors "short_url/pkg/features/shared/handlers"
	"short_url/pkg/features/shared/types"
)

var pathQuery = "sql/queries/select_users.sql"

func (pr *PostgresRepository) GetUsers(ctx context.Context) (*[]types.DbUser, error) {
	data, err := selectUsers.ReadFile(pathQuery)

	rows, err := pr.Db.QueryContext(ctx, string(data))
	if err != nil {
		return nil, &errors.BadRequestError{
			Details: "Error insterting data on database",
		}
	}

	fmt.Println(string(data))

	defer rows.Close()

	var users []types.DbUser
	for rows.Next() {
		var user types.DbUser
		if err := rows.Scan(&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Email,
			&user.Picture,
			&user.Name,
			&user.Token,
			&user.ExpiresAt); err != nil {
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
