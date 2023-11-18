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
	}
)

func (uu *authUsecase) Register(ctx context.Context, user *model.User) (*dto.UserRes, error) {
	res, err := uu.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, helper.ErrDuplicateKey) {
			return nil, helper.ErrDuplicateUser
		}
		return nil, helper.ErrInternalServer
	}
	return dto.ConvUserToRes(res), nil
}

func (uu *authUsecase) Login(ctx context.Context, user *model.User) (*dto.LoginToken, error) {
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

	aToken, err := uu.jwt.GenerateAccessToken(dto.UserTokenDTO{ID: res.ID}, uu.config)
	if err != nil {
		return nil, helper.ErrGenerateToken
	}

	rToken, err := uu.jwt.GenerateRefreshToken(uu.config)
	if err != nil {
		return nil, helper.ErrGenerateToken
	}

	return &dto.LoginToken{AToken: *aToken, RToken: *rToken}, nil
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
