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
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	FindUserByIdentifier(ctx context.Context, email string) (*model.User, error)
}

func NewUserRepo(db *gorm.DB) *usereRepo {
	return &usereRepo{
		db: db,
	}
}

func (ur *usereRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := ur.db.WithContext(ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, helper.ErrDuplicateKey
		}
		return nil, err
	}
	return user, nil
}

func (ur *usereRepo) FindUserByIdentifier(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
