package routes

import (
	auth "short_url/pkg/features/auth/graph"
	uh "short_url/pkg/features/auth/handlers"

	"github.com/labstack/echo/v4"

	dbs "short_url/pkg/databases"
)

func handlerAuthRoute(e *echo.Echo, dbs *dbs.Databases) {
	authUc := auth.WireAuthUsecases(dbs.SqlDB)

	e.GET("/auth", uh.HandlerAuth(authUc))
	e.GET("/auth/callback", uh.HandlerAuthCallback(authUc))
	e.GET("/auth/logout", uh.HandlerLogout(authUc))

}
