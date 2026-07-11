package auth


type AuthHandler struct{}

func NewAuthHandler(authService *AuthService) *AuthHandler {
    return &AuthHandler{}
}