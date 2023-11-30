package repository

import (
	"context"
	"errors"
	"mini-socmed/internal/model"
	"mini-socmed/internal/shared/errmsg"

	"gorm.io/gorm"
)

type (
	userRepo struct {
		db *gorm.DB
	}
	UserRepo interface {
		CreateUser(ctx context.Context, user *model.User) (*model.User, error)
		FindUserByIdentifier(ctx context.Context, email string) (*model.User, error)
		FirstUserByID(ctx context.Context, userID int64) (*model.User, error)
		UpdateVerifiedEmail(ctx context.Context, userID int64) (*model.User, error)
	}
)

func (r *userRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errmsg.ErrDuplicateKey
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FindUserByIdentifier(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errmsg.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) UpdateVerifiedEmail(ctx context.Context, userID int64) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Model(user).Where("email_is_verified = FALSE").Update("email_is_verified", "TRUE").Error; err != nil {
		if errors.Is(err, gorm.ErrNotImplemented) {
			return nil, errmsg.ErrUserVerifyEmail
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FirstUserByID(ctx context.Context, userID int64) (*model.User, error) {
	user := new(model.User)
	if err := r.db.WithContext(ctx).First(user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errmsg.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
