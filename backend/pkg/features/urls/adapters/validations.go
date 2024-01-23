package adapters

import (
	"fmt"
	"regexp"
	errors "short_url/pkg/features/shared/handlers"

	"github.com/labstack/echo/v4"
)

func validateUrl(ou string) (*OriginalUrl, error) {
	urlPattern := `^https?://([a-zA-Z0-9.-]+)(:[0-9]+)?/?([^\s]*)$`

	regex, err := regexp.Compile(urlPattern)
	if err != nil {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Original url is not a valid url: %v", err),
		}
	}

	match := regex.MatchString(ou)
	if !match {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Original url is not a valid url: %v", err),
		}
	}

	return &OriginalUrl{
		Ou: &ou,
	}, nil
}

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
