package routes

import (
	urls "short_url/pkg/features/urls/graph"
	uh "short_url/pkg/features/urls/handlers"

	"github.com/labstack/echo/v4"

	dbs "short_url/pkg/databases"
)

func handlerUrlsRoute(e *echo.Echo, dbs *dbs.Databases) {
	urlUc := urls.WireUrlsUsecases(dbs.SqlDB)

	e.GET("/url", uh.HandlerUrlRedirect(urlUc))
	e.POST("/url", uh.HandlerShortUrl(urlUc))
	e.DELETE("/url", uh.HandlerDeleteUrl(urlUc))
	e.POST("/url/use", uh.HandlerUrlUsage(urlUc))
	e.POST("/url/title", uh.HandlerUrlTitle(urlUc))
}
