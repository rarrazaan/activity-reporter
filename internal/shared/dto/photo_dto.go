package dto

type PhotoReq struct {
	ImageUrl string `json:"image_url" binding:"required"`
	Caption  string `json:"caption"`
}

type PhotoRes struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}
