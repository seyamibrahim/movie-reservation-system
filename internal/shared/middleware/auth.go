package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/seyamibrahim/movie-reservation-system/configs"
	sharedjwt "github.com/seyamibrahim/movie-reservation-system/internal/shared/jwt"
	"github.com/seyamibrahim/movie-reservation-system/internal/shared/response"
)

func AuthMiddleware(cfg *configs.AppConfig) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "authorization header is required")
			c.Abort()
			return
		}

		// Validate Bearer token format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "invalid authorization header")
			c.Abort()
			return
		}

		// Extract JWT
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify JWT
		claims, err := sharedjwt.VerifyJWT(token, cfg)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Store user information in request context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}