package validations

import (
	"context"
	"fmt"
	"short_url/pkg/features/shared/types"
	"short_url/pkg/features/shared/utils"

	handlers "short_url/pkg/features/shared/handlers"
)

func CheckUserIsValidated(ctx context.Context, dbUser *types.DbUser, queryId, token string) (*types.DbUser, error) {
	for i, char := range token {
		if string(dbUser.Token[i]) != string(byte(char)) {
			return nil, &handlers.UnauthorizedError{
				Details: "User token not valid",
			}
		}
	}

	err := utils.IsDateExpired(dbUser.ExpiresAt)
	if err != nil {
		return nil, &handlers.UnauthorizedError{
			Details: fmt.Sprintf("Token expired: %v", err),
		}
	}

	return dbUser, nil
}
