package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) HandleGetUserById(c echo.Context) error {
	user := s.store.Get("10")

	return c.JSON(http.StatusOK, user)
}
