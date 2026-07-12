package handlers

import (

	"github.com/seyamibrahim/movie-reservation-system/internal/auth"
	"github.com/seyamibrahim/movie-reservation-system/internal/user"
)

type Handlers struct {
	AuthHandler *auth.AuthHandler
	UserHandler *user.UserHanlder
}

func NewHandlers(authHandler *auth.AuthHandler, userHandler *user.UserHanlder) *Handlers {
	return &Handlers{
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}
