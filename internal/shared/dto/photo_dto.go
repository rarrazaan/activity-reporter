package dto

import (
	"mini-socmed/internal/model"
)

type PhotoReq struct {
	ImageUrl string `json:"image_url" binding:"required"`
	Caption  string `json:"caption"`
}

type PhotoRes struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

func ConvPhotoReq(photo *PhotoReq, userID int64) *model.Photo {
	return &model.Photo{
		ImageUrl: photo.ImageUrl,
		Caption:  photo.Caption,
		UserID:   userID,
	}
}

func ConvPhotoRes(photo *model.Photo) *PhotoRes {
	return &PhotoRes{
		ID:     photo.ID,
		UserID: photo.UserID,
	}
}
