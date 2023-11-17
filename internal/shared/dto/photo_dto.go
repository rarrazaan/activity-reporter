package dto

import (
	"activity-reporter/model"
	"io"
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

func ConvPhotoReq(photo PhotoReq, userID int64) model.Photo {
	fileContent, _ := photo.Image.Open()
	byteContainer, _ := io.ReadAll(fileContent)
	return model.Photo{
		Image:   byteContainer,
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
