package repository

import (
	"context"
	"mini-socmed/internal/model"

	"gorm.io/gorm"
)

type (
	photoRepo struct {
		db *gorm.DB
	}
	PhotoRepo interface {
		AddPhoto(ctx context.Context, photo *model.Photo) (*model.Photo, error)
	}
)

func (pr *photoRepo) AddPhoto(ctx context.Context, photo *model.Photo) (*model.Photo, error) {
	err := pr.db.WithContext(ctx).Create(photo).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &photoRepo{
		db: db,
	}
}
