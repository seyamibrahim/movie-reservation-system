package handlers

import "github.com/seyamibrahim/movie-reservation-system/internal/auth"

type Handlers struct {
	AuthHandler *auth.AuthHandler
}

func NewHandlers(authHandler *auth.AuthHandler) *Handlers {
	return &Handlers{
		AuthHandler: authHandler,
	}
}
