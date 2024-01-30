package seed

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func HandlerSeed(uc *SeedUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		ctx := c.Request().Context()

		err := uc.createSeed(ctx)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)

			return json.NewEncoder(w).Encode(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		return json.NewEncoder(w).Encode("Tables created successfully")
	}
}
