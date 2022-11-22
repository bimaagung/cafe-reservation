package mock

import (
	"context"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Create(ctx context.Context, user userdomain.User) {
	args := m.Called(ctx, user)

	var r0 error
	
	if rf, ok := args.Get(0).(func(context.Context, userdomain.User)); ok {
		r0 = rf(ctx, user)
	}else {
		r0 = args.Error(0)
	}

	return r0
}