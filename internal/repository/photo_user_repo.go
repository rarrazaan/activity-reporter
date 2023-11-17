package repository

import (
	"activity-reporter/model"
	"context"

	"gorm.io/gorm"
)

type userPhotoRepo struct {
	db *gorm.DB
}

type UserPhotoRepo interface {
	AddPhoto(ctx context.Context, photo *model.UserPhoto) (*model.UserPhoto, error)
}

func NewUserPhotoRepo(db *gorm.DB) *userPhotoRepo {
	return &userPhotoRepo{
		db: db,
	}
}

func (pr *userPhotoRepo) AddLiker(ctx context.Context, liker *model.UserPhoto) (*model.UserPhoto, error) {
	err := pr.db.WithContext(ctx).Create(&liker).Error
	if err != nil {
		return nil, err
	}
	return liker, nil
}
