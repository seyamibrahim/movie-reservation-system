package user

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin Role = "ADMIN"
	User  Role = "USER"
)

type UserModel struct {
	ID           uuid.UUID    `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(100),not null"`
	Email        string    `gorm:"type:varchar(255),not null"`
	PasswordHash string    `gorm:"type:varchar(255),not null"`
	Role         Role      `gorm:"type:varchar(20),default:USER,not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}


