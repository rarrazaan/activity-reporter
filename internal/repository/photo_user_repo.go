package repository

import (
	"context"
	"mini-socmed/internal/model"

	"gorm.io/gorm"
)

type (
	userPhotoRepo struct {
		db *gorm.DB
	}
	UserPhotoRepo interface {
		AddLiker(ctx context.Context, liker *model.UserPhoto) (*model.UserPhoto, error)
	}
)

func (pr *userPhotoRepo) AddLiker(ctx context.Context, liker *model.UserPhoto) (*model.UserPhoto, error) {
	err := pr.db.WithContext(ctx).Create(&liker).Error
	if err != nil {
		return nil, err
	}
	return liker, nil
}

func NewUserPhotoRepo(db *gorm.DB) UserPhotoRepo {
	return &userPhotoRepo{
		db: db,
	}
}
