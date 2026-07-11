package handlers

import (
	"github.com/seyamibrahim/movie-reservation-system/internal/auth"
	"gorm.io/gorm"
)

func InitHandlers(db *gorm.DB) *Handlers {

	// Initialize Repositories
	authRepo := auth.NewAuthRepository(db)

	// Initialize Services

	authService := auth.NewAuthService(authRepo)

	// Initialize Handlers

	authHandler := auth.NewAuthHandler(authService)

	return NewHandlers(authHandler)
}
