package repository

import (
	"context"
	"errors"
	"mini-socmed/internal/model"
	"mini-socmed/internal/shared/helper"

	"gorm.io/gorm"
)

type (
	usereRepo struct {
		db *gorm.DB
	}
	UserRepo interface {
		CreateUser(ctx context.Context, user *model.User) (*model.User, error)
		FindUserByIdentifier(ctx context.Context, email string) (*model.User, error)
	}
)

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

func NewUserRepo(db *gorm.DB) UserRepo {
	return &usereRepo{
		db: db,
	}
}