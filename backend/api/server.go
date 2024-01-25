// api/server.go
package api

import (
	"short_url/storage"

	"github.com/labstack/echo/v4"
)

type Server struct {
	router *echo.Echo
	store  storage.Storage
}

func NewServer(store storage.Storage) *Server {
	return &Server{
		router: echo.New(),
		store:  store,
	}
}

func (s *Server) Start(listenAddr string) error {
	s.setupRoutes()

	return s.router.Start(listenAddr)
}
