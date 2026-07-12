package server

import "github.com/gin-gonic/gin"

func SetupAuthRoutes(r *gin.RouterGroup, s *HTTPServer) {
	r.POST("/signup", s.Handlers.AuthHandler.SignUp)

	r.POST("/login", s.Handlers.AuthHandler.Login)
}
