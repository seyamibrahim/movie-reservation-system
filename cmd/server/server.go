package server

import (
	"log"

	"github.com/gin-gonic/gin"
	configs "github.com/seyamibrahim/movie-reservation-system/configs"
	h "github.com/seyamibrahim/movie-reservation-system/internal/handlers"
)

const BasePath = "/api/v1"

// HTTPServer represents the HTTP server
type HTTPServer struct {
	Engine   *gin.Engine
	Config   *configs.AppConfig
	Handlers *h.Handlers
}

func NewHTTPServer(cfg *configs.AppConfig, handlers *h.Handlers) (*HTTPServer, error) {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin engine
	engine := gin.New()

	// Set global middleware
	engine.Use(gin.Logger(), gin.Recovery())

	// Create a new HTTP server
	server := &HTTPServer{
		Engine:   engine,
		Config:   cfg,
		Handlers: handlers,
	}

	// Register routes
	server.SetupRoutes()

	return server, nil
}

func (s *HTTPServer) Run() error {
	log.Printf("Server started on :%s", s.Config.Port)
	return s.Engine.Run(":" + s.Config.Port)
}

func (s *HTTPServer) SetupRoutes() {
	SetupPublicRoutes(s)
	SetupPrivateRoutes(s)
}
