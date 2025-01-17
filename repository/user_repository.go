package repository

import (
	"context"
	"go-booking/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
		CheckEmail(ctx context.Context, email string) (entity.User, error)
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	tx := u.db

	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) CheckEmail(ctx context.Context, email string) (entity.User, error) {
	tx := u.db
	var user entity.User
	err := tx.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
