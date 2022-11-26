package mocks

import (
	"context"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/stretchr/testify/mock"
)

type MenuRepositoryPostgresMock struct {
	mock.Mock
}

func (m *MenuRepositoryPostgresMock) Add(ctx context.Context, menu *domain.Menu)(string, error){
	ret := m.Called(ctx, menu)

	var r0 string

	if rf, ok := ret.Get(0).(func(context.Context, *domain.Menu) string); ok {
		r0 = rf(ctx, menu)
	}else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Menu) error); ok {
		r1 = rf(ctx, menu)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MenuRepositoryPostgresMock) GetList(ctx context.Context)([]domain.Menu, error){
	ret := m.Called(ctx)

	var r0 []domain.Menu

	if rf, ok := ret.Get(0).(func(context.Context) []domain.Menu); ok {
		r0 = rf(ctx)
	}else {
		r0 = ret.Get(0).([]domain.Menu)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MenuRepositoryPostgresMock) Delete(ctx context.Context, id string)error{
	ret := m.Called(ctx, id)

	var r0 error

	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}

func (m *MenuRepositoryPostgresMock) GetByName(ctx context.Context, name string)(domain.Menu, error) {
	ret := m.Called(ctx, name)

	var r0 domain.Menu

	if rf, ok := ret.Get(0).(func(context.Context, string)domain.Menu); ok {
		r0 = rf(ctx, name)
	}else{
		r0 = ret.Get(0).(domain.Menu)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, string)error); ok {
		r1 = rf(ctx, name)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MenuRepositoryPostgresMock) GetById(ctx context.Context, id string)(domain.Menu, error) {
	ret := m.Called(ctx, id)

	var r0 domain.Menu

	if rf, ok := ret.Get(0).(func(context.Context, string)domain.Menu); ok {
		r0 = rf(ctx, id)
	}else{
		r0 = ret.Get(0).(domain.Menu)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, string)error); ok {
		r1 = rf(ctx, id)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MenuRepositoryPostgresMock) Update (ctx context.Context, id string, menu *domain.Menu)(string, error){
	ret := m.Called(ctx, id,menu)

	var r0 string

	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.Menu) string); ok {
		r0 = rf(ctx, id, menu)
	}else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(context.Context, string, *domain.Menu) error); ok {
		r1 = rf(ctx, id, menu)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}