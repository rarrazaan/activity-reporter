package usecase

import (
	"activity-reporter/model"
	"activity-reporter/repository"
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	crypto   helper.AppCrypto
	jwt      helper.JwtTokenizer
	userRepo repository.UserRepo
}

type UserUsecase interface {
	Register(ctx context.Context, user *model.User) (*dto.UserRes, error)
	Login(ctx context.Context, user *model.User) (*dto.LoginRes, error)
}

func NewUserUsecase(userRepo repository.UserRepo, crypto helper.AppCrypto, jwt helper.JwtTokenizer) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		crypto:   crypto,
		jwt:      jwt,
	}
}

func (uu *userUsecase) Register(ctx context.Context, user *model.User) (*dto.UserRes, error) {
	res, err := uu.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, helper.ErrDuplicateKey) {
			return nil, helper.ErrDuplicateUser
		}
		return nil, helper.ErrInternalServer
	}
	return dto.ConvUserToRes(res), nil
}

func (uu *userUsecase) Login(ctx context.Context, user *model.User) (*dto.LoginRes, error) {
	res, err := uu.userRepo.FindUserByIdentifier(ctx, user.Email)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotFound) {
			return nil, helper.ErrCredential
		}
		return nil, helper.ErrInternalServer
	}

	err = uu.crypto.ComparePasswords(user.Password, res.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, helper.ErrCredential
	}

	token, err := uu.jwt.GenerateToken(dto.UserTokenDTO{ID: res.ID})
	if err != nil {
		return nil, helper.ErrGenerateToken
	}
	return &dto.LoginRes{User: dto.ConvUserToRes(res), Token: token}, nil
}
