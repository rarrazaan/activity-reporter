package usecase

import (
	"context"
	"errors"
	"fmt"
	"mini-socmed/internal/cons"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/model"
	"mini-socmed/internal/repository"
	"mini-socmed/internal/shared/dto"
	"mini-socmed/internal/shared/errmsg"
	"mini-socmed/internal/shared/helper"

	"golang.org/x/crypto/bcrypt"
)

type (
	authUsecase struct {
		crypto    helper.AppCrypto
		jwt       helper.JwtTokenizer
		uuid      helper.UuidGenerator
		config    dependency.Config
		userRepo  repository.UserRepo
		redisRepo repository.RedisRepo
		euc       EmailSenderUsecase
	}

	AuthUsecase interface {
		Register(ctx context.Context, user *model.User) (*dto.UserRes, error)
		Login(ctx context.Context, user *model.User) (*dto.LoginToken, error)
		RefreshAccessToken(ctx context.Context, rToken string) (*string, error)
		VerifyEmail(ctx context.Context, code string) error
	}
)

func (au *authUsecase) Register(ctx context.Context, user *model.User) (*dto.UserRes, error) {
	res, err := au.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, errmsg.ErrDuplicateKey) {
			return nil, errmsg.ErrDuplicateUser
		}
		return nil, err
	}
	code := au.uuid.GenerateUUID()
	if err := au.redisRepo.SetVerifiyEmail(ctx, code, res.ID); err != nil {
		return nil, err
	}
	content := fmt.Sprintf(cons.VerificationEmailContent, "http://localhost:8080/auth", code)
	if err := au.euc.SendEmail(cons.SubjectVerificationEmail, content, user.Email); err != nil {
		return nil, err
	}
	return dto.ConvUserToRes(res), nil
}

func (au *authUsecase) Login(ctx context.Context, user *model.User) (*dto.LoginToken, error) {
	res, err := au.userRepo.FindUserByIdentifier(ctx, user.Email)
	if err != nil {
		if errors.Is(err, errmsg.ErrUserNotFound) {
			return nil, errmsg.ErrCredential
		}
		return nil, err
	}

	err = au.crypto.ComparePasswords(user.Password, res.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errmsg.ErrCredential
	}

	aToken, err := au.jwt.GenerateAccessToken(dto.UserTokenDTO{ID: res.ID}, au.config)
	if err != nil {
		return nil, errmsg.ErrGenerateToken
	}

	rToken, err := au.jwt.GenerateRefreshToken(au.config)
	if err != nil {
		return nil, errmsg.ErrGenerateToken
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

func (au *authUsecase) VerifyEmail(ctx context.Context, code string) error {
	userID, err := au.redisRepo.GetUserIDByVerifiyEmail(ctx, code)
	if err != nil {
		return err
	}
	if _, err := au.userRepo.UpdateVerifiedEmail(ctx, *userID); err != nil {
		return err
	}
	if err := au.redisRepo.DeleteVerifiyEmail(ctx, code); err != nil {
		return err
	}
	return nil
}

func NewUserUsecase(
	userRepo repository.UserRepo,
	redisRepo repository.RedisRepo,
	crypto helper.AppCrypto,
	jwt helper.JwtTokenizer,
	config dependency.Config,
	uuid helper.UuidGenerator,
	euc EmailSenderUsecase,
) AuthUsecase {
	return &authUsecase{
		redisRepo: redisRepo,
		userRepo:  userRepo,
		crypto:    crypto,
		jwt:       jwt,
		config:    config,
		uuid:      uuid,
		euc:       euc,
	}
}
