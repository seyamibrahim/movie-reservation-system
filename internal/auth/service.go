package auth

import (
	"context"

	"github.com/seyamibrahim/movie-reservation-system/configs"
	"github.com/seyamibrahim/movie-reservation-system/internal/shared/jwt"
	"github.com/seyamibrahim/movie-reservation-system/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo user.IUserRepository
	config   *configs.AppConfig
}

func NewAuthService(userRepo user.IUserRepository, config *configs.AppConfig) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		config:   config,
	}
}

func (a *AuthService) SignUp(ctx context.Context, req *SignUpRequest) (*AuthResponse, error) {

	// Check if email already exists
	exists, err := a.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}
	// Create user
	newUser := &user.UserModel{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         user.User,
	}

	if err := a.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	// Generate JWT
	accessToken, err := jwt.CreateJWT(
		newUser.ID,
		newUser.Email,
		string(newUser.Role),
		a.config,
	)

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User: UserResponse{
			ID:    newUser.ID,
			Name:  newUser.Name,
			Email: newUser.Email,
			Role:  string(newUser.Role),
		},
	}, nil
}

// Login authenticates an existing user.
func (a *AuthService) Login(
	ctx context.Context,
	req *LoginRequest,
) (*AuthResponse, error) {

	// Find user
	existingUser, err := a.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(existingUser.PasswordHash),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT
	accessToken, err := jwt.CreateJWT(
		existingUser.ID,
		existingUser.Email,
		string(existingUser.Role),
		a.config,
	)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User: UserResponse{
			ID:    existingUser.ID,
			Name:  existingUser.Name,
			Email: existingUser.Email,
			Role:  string(existingUser.Role),
		},
	}, nil
}
