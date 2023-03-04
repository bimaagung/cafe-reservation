package mocks

import (
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type MenuMinioRepository struct {
	mock.Mock
}

func (m *MenuMinioRepository) Upload(file *multipart.FileHeader, bucketName string, objectName string) error{
	ret := m.Called(file, bucketName, objectName)

	var r0 error

	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, string, string) error); ok {
		r0 = rf(file, bucketName, objectName)
	}else{
		r0 = ret.Error(0)
	}

	return r0
}
