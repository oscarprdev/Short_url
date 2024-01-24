package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
	urlUc "short_url/pkg/features/urls/usecases"

	"github.com/labstack/echo/v4"
)

func provideUrlBody(c echo.Context, w http.ResponseWriter) (*api.PostUrlTitleJSONBody, error) {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		newError := &errors.BadRequestError{
			Details: fmt.Sprintf("Error reading request body: %v", err),
		}
		return nil, handleUrlsErrors(w, newError)
	}

	var rb api.PostUrlTitleJSONBody
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

	return &api.PostUrlTitleJSONBody{
		UrlId:    rb.UrlId,
		UrlTitle: rb.UrlTitle,
	}, nil
}

func HandlerUrlTitle(uc *urlUc.UrlUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		ctx := c.Request().Context()

		urlBody, err := provideUrlBody(c, w)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Not valid request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		authToken := c.Request().Header.Get("Authorization")
		idQueryParam := c.QueryParam("id")

		// Initialization of url response
		url, err := uc.UrlTitleUsecases(ctx, urlBody, idQueryParam, authToken)
		if err != nil {
			return handleUrlsErrors(w, err)
		}

		// Return successful response
		status := int(200)
		successResponse := api.PostUrlTitle200JSONResponse{
			Status: &status,
			Url:    url,
		}

		return successResponse.VisitPostUrlTitleResponse(w)
	}
}
