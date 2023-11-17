// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	dto "mini-socmed/internal/shared/dto"
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "mini-socmed/internal/model"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GoogleLogin provides a mock function with given fields: ctx, user
func (_m *UserUsecase) GoogleLogin(ctx context.Context, user model.User) (dto.LoginRes, error) {
	ret := _m.Called(ctx, user)

	var r0 dto.LoginRes
	if rf, ok := ret.Get(0).(func(context.Context, model.User) dto.LoginRes); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(dto.LoginRes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, user
func (_m *UserUsecase) Login(ctx context.Context, user *model.User) (*dto.LoginRes, error) {
	ret := _m.Called(ctx, user)

	var r0 *dto.LoginRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*dto.LoginRes, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *dto.LoginRes); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, user
func (_m *UserUsecase) Register(ctx context.Context, user *model.User) (*dto.UserRes, error) {
	ret := _m.Called(ctx, user)

	var r0 *dto.UserRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*dto.UserRes, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *dto.UserRes); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}