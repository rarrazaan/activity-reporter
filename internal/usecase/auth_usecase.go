package usecase

import (
	"context"
	"errors"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/model"
	"mini-socmed/internal/repository"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/shared/helper"

	"golang.org/x/crypto/bcrypt"
)

type (
	authUsecase struct {
		crypto    helper.AppCrypto
		jwt       helper.JwtTokenizer
		config    dependency.Config
		userRepo  repository.UserRepo
		redisRepo repository.RedisRepo
	}

	AuthUsecase interface {
		Register(ctx context.Context, user *model.User) (*dto.UserRes, error)
		Login(ctx context.Context, user *model.User) (*dto.LoginToken, error)
		RefreshAccessToken(ctx context.Context, rToken string) (*string, error)
	}
)

func (au *authUsecase) Register(ctx context.Context, user *model.User) (*dto.UserRes, error) {
	res, err := au.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, helper.ErrDuplicateKey) {
			return nil, helper.ErrDuplicateUser
		}
		return nil, err
	}
	return dto.ConvUserToRes(res), nil
}

func (au *authUsecase) Login(ctx context.Context, user *model.User) (*dto.LoginToken, error) {
	res, err := au.userRepo.FindUserByIdentifier(ctx, user.Email)
	if err != nil {
		if errors.Is(err, helper.ErrUserNotFound) {
			return nil, helper.ErrCredential
		}
		return nil, err
	}

	err = au.crypto.ComparePasswords(user.Password, res.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, helper.ErrCredential
	}

	aToken, err := au.jwt.GenerateAccessToken(dto.UserTokenDTO{ID: res.ID}, au.config)
	if err != nil {
		return nil, helper.ErrGenerateToken
	}

	rToken, err := au.jwt.GenerateRefreshToken(au.config)
	if err != nil {
		return nil, helper.ErrGenerateToken
	}
	if err := au.redisRepo.SetRefreshToken(ctx, *rToken, res.ID); err != nil {
		return nil, err
	}

	return &dto.LoginToken{AToken: *aToken, RToken: *rToken}, nil
}

func (au *authUsecase) RefreshAccessToken(ctx context.Context, rToken string) (*string, error) {
	_, err := au.jwt.ValidateRefreshToken(rToken, au.config)
	if err != nil {
		return nil, err
	}
	userID, err := au.redisRepo.GetUserIDByRefreshToken(ctx, rToken)
	if err != nil {
		return nil, err
	}
	aToken, err := au.jwt.GenerateAccessToken(dto.UserTokenDTO{ID: *userID}, au.config)
	if err != nil {
		return nil, err
	}
	return aToken, nil
}

func NewUserUsecase(
	userRepo repository.UserRepo,
	redisRepo repository.RedisRepo,
	crypto helper.AppCrypto,
	jwt helper.JwtTokenizer,
	config dependency.Config,
) AuthUsecase {
	return &authUsecase{
		redisRepo: redisRepo,
		userRepo:  userRepo,
		crypto:    crypto,
		jwt:       jwt,
		config:    config,
	}
}
