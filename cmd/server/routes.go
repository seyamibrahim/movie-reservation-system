package server

import "github.com/seyamibrahim/movie-reservation-system/internal/shared/middleware"

func SetupPublicRoutes(s *HTTPServer) {
	srs := s.Engine.Group(BasePath)

	authGroup := srs.Group("/auth")

	SetupAuthRoutes(authGroup, s)

}

func SetupPrivateRoutes(s *HTTPServer) {

	srs := s.Engine.Group(BasePath)

	srs.Use(middleware.AuthMiddleware(s.Config))


	// protected routes
	userGroup := srs.Group("/user")

	SetupUserRoutes(userGroup, s)
}
