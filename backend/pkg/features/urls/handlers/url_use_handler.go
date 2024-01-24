package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
	urlUc "short_url/pkg/features/urls/usecases"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func provideUrlId(c echo.Context, w http.ResponseWriter) (*uuid.UUID, error) {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	var rb api.PostUrlUseJSONBody
	err = json.Unmarshal(b, &rb)
	if err != nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Error JSON format on request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	if rb.UrlId == nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Not valid request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	return rb.UrlId, nil
}

func HandlerUrlUsage(uc *urlUc.UrlUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		ctx := c.Request().Context()

		urlId, err := provideUrlId(c, w)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Not valid request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		authToken := c.Request().Header.Get("Authorization")
		idQueryParam := c.QueryParam("id")

		// Initialization of url response
		url, err := uc.UrlUsageUsecases(ctx, *urlId, idQueryParam, authToken)
		if err != nil {
			return handleUrlsErrors(w, err)
		}

		// Return successful response
		status := int(200)
		successResponse := api.PostUrl200JSONResponse{
			Status: &status,
			Url:    url,
		}

		return successResponse.VisitPostUrlResponse(w)
	}
}
