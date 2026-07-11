package auth

import (
    "gorm.io/gorm"
)

type AuthRepository struct {}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
    return &AuthRepository{}
}