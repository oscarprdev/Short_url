package utils

import (
	"time"

	errorsC "short_url/pkg/features/shared/handlers"
)

func IsDateExpired(t time.Time) error {
	now := time.Now().UTC()
	if now.After(t) {
		return &errorsC.UnauthorizedError{
			Details: "Token expired",
		}
	}

	return nil
}
