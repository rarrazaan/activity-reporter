package dto

import (
	"activity-reporter/model"
	"mime/multipart"
)

type PhotoReq struct {
	Image   *multipart.FileHeader `form:"image" binding:"required"`
	Caption string                `form:"caption"`
}

type PhotoRes struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

func ConvPhotoReq(photo PhotoReq, urlImg string, userID int64) model.Photo {
	return model.Photo{
		Image:   urlImg,
		Caption: photo.Caption,
		UserID:  userID,
	}
}

func ConvPhotoRes(photo model.Photo) PhotoRes {
	return PhotoRes{
		ID:     photo.ID,
		UserID: photo.UserID,
	}
}
