package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seyamibrahim/movie-reservation-system/internal/shared/response"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.authService.SignUp(c.Request.Context(), &req)
	if err != nil {

		switch err {

		case ErrEmailAlreadyExists:
			response.Error(c, http.StatusConflict, err.Error())

		default:
			response.Error(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.Created(c, "User registered successfully", res)
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {

		switch err {

		case ErrInvalidCredentials:
			response.Error(c, http.StatusUnauthorized, err.Error())

		default:
			response.Error(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.OK(c, "Login successful", res)
}
