package handlers

import (
	api "short_url/pkg/api"
	urlUc "short_url/pkg/features/urls/usecases"

	"github.com/labstack/echo/v4"
)

func HandlerDeleteUrl(uc *urlUc.UrlUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		ctx := c.Request().Context()

		userId := c.QueryParam("id")
		urlId := c.QueryParam("urlId")

		// Initialization of url response
		err := uc.RemoveUrlById(ctx, userId, urlId)
		if err != nil {
			return handleUrlsErrors(w, err)
		}

		// Return successful response
		status := int(200)
		message := "Url deleted successfully"
		successResponse := api.DeleteUrl200JSONResponse{
			Status:  &status,
			Message: &message,
		}

		return successResponse.VisitDeleteUrlResponse(w)
	}
}
