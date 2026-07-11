package server

func SetupPublicRoutes(s *HTTPServer) {
	srs := s.Engine.Group(BasePath)

	authGroup := srs.Group("/auth")

	SetupAuthRoutes(authGroup, s)

}

func SetupPrivateRoutes(s *HTTPServer) {




}
