package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
	adapters "short_url/pkg/features/urls/adapters"
	urlUc "short_url/pkg/features/urls/usecases"

	"github.com/labstack/echo/v4"
)

func provideOriginalUrl(c echo.Context, w http.ResponseWriter) (*adapters.OriginalUrl, error) {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	var rb api.PostUrlJSONBody
	err = json.Unmarshal(b, &rb)
	if err != nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Error JSON format on request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	if rb.OriginalUrl == nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Not valid request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	r := rb.OriginalUrl

	return &adapters.OriginalUrl{Ou: r}, nil
}

func HandlerShortUrl(uc *urlUc.UrlUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		ctx := c.Request().Context()

		ou, err := provideOriginalUrl(c, w)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Not valid request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		authToken := c.Request().Header.Get("Authorization")
		idQueryParam := c.QueryParam("id")

		// Initialization of url response
		var url = api.Url{}

		if idQueryParam != "" && authToken != "" {
			// Manage the short url and connect url with user to be persistent in database
			err = uc.ShortUrlComplexUsecases(ctx, ou, &url, idQueryParam, authToken)
			if err != nil {
				newError := &errors.InternalError{
					Details: fmt.Sprintf("Internal error: %v", err),
				}
				return handleUrlsErrors(w, newError)
			}
		} else {
			// Just short url and return back to http response
			err = uc.ShortUrlSimpleUsecases(ou, &url)
			if err != nil {
				newError := &errors.InternalError{
					Details: fmt.Sprintf("Internal error: %v", err),
				}
				return handleUrlsErrors(w, newError)
			}
		}

		// Return successful response
		status := int(200)
		successResponse := api.PostUrl200JSONResponse{
			Status: &status,
			Url:    &url,
		}

		return successResponse.VisitPostUrlResponse(w)
	}
}
