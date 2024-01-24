package postgres

import (
	"context"
	"fmt"
	errors "short_url/pkg/features/shared/handlers"

	"github.com/google/uuid"
)

var insertPathQuery = "sql/queries/insert_url_user.sql"

type UrlUserInput struct {
	Id     uuid.UUID `json:"id"`
	UserId string    `json:"user_id"`
	UrlId  uuid.UUID `json:"url_id"`
}

func (pr *PostgresRepository) RelateUrlWithUser(ctx context.Context, urlId uuid.UUID, userId string) error {

	query, err := insertUrlUserQuery.ReadFile(insertPathQuery)
	if err != nil {
		return &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading query: %v", err),
		}
	}

	row := pr.Db.QueryRowContext(ctx, string(query), uuid.New(), userId, urlId)

	var uui UrlUserInput
	err = row.Scan(
		&uui.Id,
		&uui.UserId,
		&uui.UrlId,
	)

	if err != nil {
		return &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading URL from database: %v", err),
		}
	}

	return nil
}
