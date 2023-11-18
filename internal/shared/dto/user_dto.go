package dto

import "mini-socmed/internal/model"

type RegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserRes struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ConvURegisToModel(req *RegisterReq) *model.User {
	return &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ConvUserToRes(user *model.User) *UserRes {
	return &UserRes{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	User  *UserRes
	Token string
}

func ConvULoginToModel(req *LoginReq) *model.User {
	return &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

type UserTokenDTO struct {
	ID int64 `json:"id"`
}
