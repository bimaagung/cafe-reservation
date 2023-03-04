package mocks

import (
	"context"

	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	ret := m.Called(ctx, username)

	var r0 domain.User

	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, string)  error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *UserRepository) Create(ctx context.Context, user *domain.User)(string, error) {
	ret := m.Called(ctx, user)

	var r0 string

	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, *domain.User)  error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}