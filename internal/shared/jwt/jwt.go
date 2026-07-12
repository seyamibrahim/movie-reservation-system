package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/seyamibrahim/movie-reservation-system/configs"
    "github.com/seyamibrahim/movie-reservation-system/internal/user"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   user.Role `json:"role"`

	jwt.RegisteredClaims
}

// CreateJWT generates a signed JWT token.
func CreateJWT(userID uuid.UUID, email, role string, cfg *configs.AppConfig) (string, error) {

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   user.Role(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.JWTSecret))
}

// VerifyJWT validates and parses a JWT token.
func VerifyJWT(tokenString string, cfg *configs.AppConfig) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		getSecretKey(cfg),
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func getSecretKey(cfg *configs.AppConfig) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidToken
		}

		return []byte(cfg.JWTSecret), nil
	}
}
