package repository

import (
	"activity-reporter/model"
	"activity-reporter/shared/helper"
	"context"
	"errors"

	"gorm.io/gorm"
)

type usereRepo struct {
	db *gorm.DB
}

type UserRepo interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	FindUserByIdentifier(ctx context.Context, email string) (model.User, error)
}

func NewUserRepo(db *gorm.DB) *usereRepo {
	return &usereRepo{
		db: db,
	}
}

func (ur *usereRepo) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	if err := ur.db.WithContext(ctx).Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return model.User{}, helper.ErrDuplicateKey
		}
		return model.User{}, err
	}
	return user, nil
}

func (ur *usereRepo) FindUserByIdentifier(ctx context.Context, email string) (model.User, error) {
	var user model.User
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, helper.ErrUserNotFound
		}
		return model.User{}, err
	}
	return user, nil
}
