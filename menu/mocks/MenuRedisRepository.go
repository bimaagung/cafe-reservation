package mocks

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/stretchr/testify/mock"
)

type MenuRedisRepository struct {
	mock.Mock
}

func (m *MenuRedisRepository) Set(menu interface{})(string, error){
	ret := m.Called(menu)

	var r0 string

	if rf, ok := ret.Get(0).(func(interface{}) string); ok {
		r0 = rf(menu)
	}else {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(menu)
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MenuRedisRepository) Get()([]domain.Menu, error){
	ret := m.Called()

	var r0 []domain.Menu

	if rf, ok := ret.Get(0).(func() []domain.Menu); ok {
		r0 = rf()
	}else {
		r0 = ret.Get(0).([]domain.Menu)
	}

	var r1 error

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	}else{
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MenuRedisRepository) Delete()error{
	ret := m.Called()

	var r0 error

	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	}else{
		r0 = ret.Error(0)
	}

	return r0
}
