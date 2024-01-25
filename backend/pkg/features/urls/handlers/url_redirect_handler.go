package handlers

import (
	"fmt"
	errors "short_url/pkg/features/shared/handlers"
	urlUc "short_url/pkg/features/urls/usecases"

	"github.com/labstack/echo/v4"
)

func HandlerUrlRedirect(uc *urlUc.UrlUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		ctx := c.Request().Context()

		url := c.QueryParam("url")

		original, err := uc.UrlRedirectUsecases(ctx, url)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Not valid request: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		// Redirect to original url
		c.Redirect(200, original)
		return nil
	}
}
