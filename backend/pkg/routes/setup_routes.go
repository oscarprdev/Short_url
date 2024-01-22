package routes

import (
	"github.com/labstack/echo/v4"

	dbs "short_url/pkg/databases"
)

func SetupRoutes(dbs *dbs.Databases) *echo.Echo {
	e := echo.New()

	handlerUsersRoute(e, dbs)
	handlerUrlsRoute(e, dbs)

	return e
}
