package usecase

import (
	"context"
	"mini-socmed/internal/model"
	"mini-socmed/internal/repository"
	"mini-socmed/internal/shared/errmsg"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	photoUsecase struct {
		photoRepo repository.PhotoRepo
		userRepo  repository.UserRepo
	}
	PhotoUsecase interface {
		PostPhoto(ctx context.Context, photo *model.Photo) (*mongo.InsertOneResult, error)
	}
)

func (pu *photoUsecase) PostPhoto(ctx context.Context, photo *model.Photo) (*mongo.InsertOneResult, error) {
	user, err := pu.userRepo.FirstUserByID(ctx, photo.UserID)
	if err != nil {
		return nil, errmsg.ErrUserNotFound
	}
	photo.Username = user.Name
	photo.CreatedAt = time.Now()
	res, err := pu.photoRepo.AddPhoto(ctx, photo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewPhotoUsecase(photoRepo repository.PhotoRepo, userRepo repository.UserRepo) PhotoUsecase {
	return &photoUsecase{
		photoRepo: photoRepo,
		userRepo:  userRepo,
	}
}
