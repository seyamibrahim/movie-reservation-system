package auth

type AuthService struct{}


func NewAuthService(aRepo *AuthRepository) *AuthService {
    return &AuthService{}
}