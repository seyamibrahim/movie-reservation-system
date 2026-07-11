package main

import (
	"log"

	"github.com/seyamibrahim/movie-reservation-system/cmd/server"
	"github.com/seyamibrahim/movie-reservation-system/configs"
	"github.com/seyamibrahim/movie-reservation-system/internal/handlers"
	"github.com/seyamibrahim/movie-reservation-system/internal/shared/database"
)

func main() {

	// Load Config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config from env: %v", err)
	}
	// Connect to Database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

    handlers := handlers.InitHandlers(db)

	// Initialize Server
	s, err := server.NewHTTPServer(cfg, handlers)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

    // Run Server
	if err := s.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }

}
