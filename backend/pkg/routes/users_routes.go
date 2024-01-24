package routes

import (
	users "short_url/pkg/features/users/graph"
	uh "short_url/pkg/features/users/handlers"

	"github.com/labstack/echo/v4"

	dbs "short_url/pkg/databases"
)

func handlerUsersRoute(e *echo.Echo, dbs *dbs.Databases) {
	userUc := users.WireUserUsecases(dbs.SqlDB)

	e.GET("/users/list", uh.HandlerGetUsers(userUc))
	e.GET("/users/describe", uh.HandlerDescribeUser(userUc))
}
