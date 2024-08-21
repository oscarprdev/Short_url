package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	dbs "short_url/pkg/databases"
)

func SetupRoutes(dbs *dbs.Databases) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "https://oplink.vercel.app"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
		MaxAge:       300,
	}))

	handlerUsersRoute(e, dbs)
	handlerUrlsRoute(e, dbs)
	handlerAuthRoute(e, dbs)
	handlerSeedRoute(e, dbs)

	return e
}
