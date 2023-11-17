package repository

import (
	"activity-reporter/model"
	"context"

	"gorm.io/gorm"
)

type photoRepo struct {
	db *gorm.DB
}

type PhotoRepo interface {
	AddPhoto(ctx context.Context, photo *model.Photo) (*model.Photo, error)
}

func NewPhotoRepo(db *gorm.DB) *photoRepo {
	return &photoRepo{
		db: db,
	}
}

func (pr *photoRepo) AddPhoto(ctx context.Context, photo *model.Photo) (*model.Photo, error) {
	err := pr.db.WithContext(ctx).Create(photo).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}
