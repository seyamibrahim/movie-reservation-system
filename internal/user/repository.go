package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user *UserModel) error
	FindByEmail(ctx context.Context, email string) (*UserModel, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Update(ctx context.Context, user *UserModel) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}


func (u *UserRepository) Create(ctx context.Context, user *UserModel) error {
    return u.db.Create(user).Error
}


func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*UserModel, error) {
    var user UserModel
    if err := u.db.First(&user, "email = ?", email).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (u *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
    var count int64
    if err := u.db.Model(&UserModel{}).Where("email = ?", email).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}


func (u *UserRepository) Update(ctx context.Context, user *UserModel) error {
    return u.db.Save(user).Error
}

func (u *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
    return u.db.Delete(&UserModel{}, "id = ?", id).Error
}