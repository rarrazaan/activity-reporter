package usecase

import (
	"activity-reporter/model"
	"activity-reporter/repository"
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"context"
)

type photoUsecase struct {
	photoRepo repository.PhotoRepo
}

type PhotoUsecase interface {
	PostPhoto(ctx context.Context, photo *model.Photo) (*dto.PhotoRes, error)
}

func NewPhotoUsecase(photoRepo repository.PhotoRepo) *photoUsecase {
	return &photoUsecase{
		photoRepo: photoRepo,
	}
}

func (pu *photoUsecase) PostPhoto(ctx context.Context, photo *model.Photo) (*dto.PhotoRes, error) {
	res, err := pu.photoRepo.AddPhoto(ctx, photo)
	if err != nil {
		return nil, helper.ErrInternalServer
	}
	return dto.ConvPhotoRes(res), nil
}
