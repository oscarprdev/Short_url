package urls

import (
	"encoding/json"
	"io"
	api "short_url/pkg/api"
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

		var rb api.PostUrlJSONBody
		if err = json.Unmarshal(b, &rb); err != nil {
			return handleUrlsErrors(w, err)
		}

		shortUrl := urlUc.ShortUrlUsecases(*rb.OriginalUrl)

		// Adapt request body to insert in DB
		rbok, err := adapters.AdaptShortUrlToDB(*rb.OriginalUrl, shortUrl)
		if err != nil {
			return handleUrlsErrors(w, err)
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
