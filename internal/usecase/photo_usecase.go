package usecase

import (
	"context"
	"mini-socmed/internal/model"
	"mini-socmed/internal/repository"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/shared/helper"
)

type (
	photoUsecase struct {
		photoRepo repository.PhotoRepo
	}
	PhotoUsecase interface {
		PostPhoto(ctx context.Context, photo *model.Photo) (*dto.PhotoRes, error)
	}
)

func (pu *photoUsecase) PostPhoto(ctx context.Context, photo *model.Photo) (*dto.PhotoRes, error) {
	res, err := pu.photoRepo.AddPhoto(ctx, photo)
	if err != nil {
		return nil, helper.ErrInternalServer
	}
	return dto.ConvPhotoRes(res), nil
}

func NewPhotoUsecase(photoRepo repository.PhotoRepo) PhotoUsecase {
	return &photoUsecase{
		photoRepo: photoRepo,
	}
}
