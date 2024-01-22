package urls

import (
	"encoding/json"
	"fmt"
	"io"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
	adapters "short_url/pkg/features/urls/adapters"
	urlUc "short_url/pkg/features/urls/usecases"
	"time"

	"github.com/labstack/echo/v4"
)

func HandlerShortUrl(urlUc *urlUc.UrlUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer

		// Provide request body
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Error reading request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		var rb api.PostUrlJSONBody
		err = json.Unmarshal(b, &rb)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Error JSON format on request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		if rb.OriginalUrl == nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Not valid request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		shortUrl := urlUc.ShortUrlUsecases(*rb.OriginalUrl)

		// Adapt request body to insert in DB
		rbok, err := adapters.AdaptShortUrlToDB(*rb.OriginalUrl, shortUrl)
		if err != nil {
			newError := &errors.BadRequestError{
				Details: fmt.Sprintf("Not valid request body: %v", err),
			}
			return handleUrlsErrors(w, newError)
		}

		// Create new response url
		now := time.Now().UTC()
		title := "Default url title"
		url := api.Url{
			OriginalUrl: &rbok.OriginalUrl,
			ShortUrl:    &shortUrl,
			UpdatedAt:   &now,
			CreatedAt:   &now,
			ExpiresAt:   &now,
			TitleUrl:    &title,
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
