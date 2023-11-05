package usecase_test

import (
	"activity-reporter/mocks"
	"activity-reporter/model"
	"activity-reporter/shared/dto"
	"activity-reporter/shared/helper"
	"activity-reporter/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var (
	mockUserRepo *mocks.UserRepo
	crypto       *mocks.AppCrypto
	jwt          *mocks.JwtTokenizer
)

func Setup() {
	mockUserRepo = new(mocks.UserRepo)
	crypto = new(mocks.AppCrypto)
	jwt = new(mocks.JwtTokenizer)
}

func Test_userUsecase_Register(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return added user when successfully register", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}
		userRegistered := &model.User{
			ID:       1,
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}

		mockUserRepo.On("CreateUser", mock.Anything, user).Return(userRegistered, nil)
		expected := &dto.UserRes{
			ID:    1,
			Name:  "Chitanda",
			Email: "chitanda@example.com",
		}

		res, err := uu.Register(context.Background(), user)

		assert.Equal(expected, res)
		assert.Nil(err)
	})

	t.Run("should return error when failed to create user", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}

		mockUserRepo.On("CreateUser", mock.Anything, user).Return(nil, errors.New("error"))

		res, err := uu.Register(context.Background(), user)

		assert.Nil(res)
		assert.ErrorIs(err, helper.ErrInternalServer)
	})
}

func Test_userUsecase_Login(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return correct user when successfully login", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}
		userExist := &model.User{
			ID:       1,
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}

		mockUserRepo.On("FindUserByIdentifier", mock.Anything, user.Email).Return(userExist, nil)
		crypto.On("ComparePasswords", user.Password, userExist.Password).Return(nil)
		jwt.On("GenerateToken", dto.UserTokenDTO{ID: userExist.ID}).Return("token", nil)
		expected := &dto.LoginRes{
			User: &dto.UserRes{
				ID:    1,
				Name:  "Chitanda",
				Email: "chitanda@example.com",
			},
			Token: "token",
		}

		res, err := uu.Login(context.Background(), user)

		assert.Equal(expected, res)
		assert.Nil(err)
	})

	t.Run("should return error when user does not exist", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password12",
		}

		mockUserRepo.On("FindUserByIdentifier", mock.Anything, user.Email).Return(nil, helper.ErrUserNotFound)

		res, err := uu.Login(context.Background(), user)

		assert.Nil(res)
		assert.ErrorIs(err, helper.ErrCredential)
	})

	t.Run("should return error when failed to find user", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password12",
		}

		mockUserRepo.On("FindUserByIdentifier", mock.Anything, user.Email).Return(nil, errors.New("error"))

		res, err := uu.Login(context.Background(), user)

		assert.Nil(res)
		assert.ErrorIs(err, helper.ErrInternalServer)
	})

	t.Run("should return error when failed to compare passwords", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password12",
		}
		userExist := &model.User{
			ID:       1,
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}

		mockUserRepo.On("FindUserByIdentifier", mock.Anything, user.Email).Return(userExist, nil)
		crypto.On("ComparePasswords", user.Password, userExist.Password).Return(bcrypt.ErrMismatchedHashAndPassword)

		res, err := uu.Login(context.Background(), user)

		assert.Nil(res)
		assert.ErrorIs(err, helper.ErrCredential)
	})

	t.Run("should return error when failed to generate JWT", func(t *testing.T) {
		Setup()
		uu := usecase.NewUserUsecase(mockUserRepo, crypto, jwt)

		user := &model.User{
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}
		userExist := &model.User{
			ID:       1,
			Name:     "Chitanda",
			Email:    "chitanda@example.com",
			Password: "password123",
		}

		mockUserRepo.On("FindUserByIdentifier", mock.Anything, user.Email).Return(userExist, nil)
		crypto.On("ComparePasswords", user.Password, userExist.Password).Return(nil)
		jwt.On("GenerateToken", dto.UserTokenDTO{ID: userExist.ID}).Return("", errors.New("error"))

		res, err := uu.Login(context.Background(), user)

		assert.Nil(res)
		assert.ErrorIs(err, helper.ErrGenerateToken)
	})
}
