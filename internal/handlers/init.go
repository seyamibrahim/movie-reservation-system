package handlers

import (
	"github.com/seyamibrahim/movie-reservation-system/configs"
	"github.com/seyamibrahim/movie-reservation-system/internal/auth"
	"github.com/seyamibrahim/movie-reservation-system/internal/user"
	"gorm.io/gorm"
)

func InitHandlers(db *gorm.DB, cfg *configs.AppConfig) *Handlers {

	// Initialize Repositories

	userRepo := user.NewUserRepository(db)

	// Initialize Services

	authService := auth.NewAuthService(userRepo, cfg)

	userService := user.NewUserService(userRepo)

	// Initialize Handlers

	authHandler := auth.NewAuthHandler(authService)

	userHandler := user.NewUserHandler(userService)

	return NewHandlers(authHandler, userHandler)
}
