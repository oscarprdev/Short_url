package api

func (s *Server) setupRoutes() {
	s.router.GET("/user", s.HandleGetUserById)
}
