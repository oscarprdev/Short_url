package adapters

import (
	errors "short_url/pkg/features/shared/handlers"

	"github.com/labstack/echo/v4"
)

func validateOrigin(c echo.Context) error {
	host := c.Request().Host
	path := c.Request().URL.Path

	if host != "http://localhost:8080" {
		return &errors.BadRequestError{
			Details: "Not valid host",
		}
	}

	if path != "/home" {
		return &errors.BadRequestError{
			Details: "Not valid path",
		}
	}

	return nil
}
