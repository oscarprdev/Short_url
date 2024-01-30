package routes

import (
	seed "short_url/pkg/features/seed"
	uh "short_url/pkg/features/seed"

	"github.com/labstack/echo/v4"

	dbs "short_url/pkg/databases"
)

func handlerSeedRoute(e *echo.Echo, dbs *dbs.Databases) {
	seedUc := seed.WireSeedUsecases(dbs.SqlDB)

	e.POST("/seed", uh.HandlerSeed(seedUc))
}
