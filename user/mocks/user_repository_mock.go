package mocks

import (
	"context"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) GetByUsername(ctx context.Context, username string) (userdomain.User, error) {
	ret := m.Called(ctx, username)

	var r0 userdomain.User

	if rf, ok := ret.Get(0).(func(context.Context, string) userdomain.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(userdomain.User)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, string)  error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *UserRepository) Create(ctx context.Context, user userdomain.User)(string, error) {
	ret := m.Called(ctx, user)

	var r0 string

	if rf, ok := ret.Get(0).(func(context.Context, userdomain.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, userdomain.User)  error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}