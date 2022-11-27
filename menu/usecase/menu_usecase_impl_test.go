package usecase_test

import (
	"context"
	"mime/multipart"
	"net/textproto"
	"testing"
	"time"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/menu/mocks"
	"github.com/bimaagung/cafe-reservation/menu/usecase"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const patchEnv = "../../.env"

func TestAddMenuUseCase(t *testing.T) {
	godotenv.Load(patchEnv)

	mockMenuRepoPostgres := new(mocks.MenuRepositoryPostgresMock) 
	mockMenuRepoRedis := new(mocks.MenuRepositoryRedisMock) 
	mockMenuRepoMinio := new(mocks.MenuRepositoryMinioMock) 
	
	mockMenuReq := domain.MenuReq{
		Id: "97391bbb-a48f-48e2-a166-db669e6377fc",
		Name: "Cappucino",
		Price: 40000,
		Stock: 100,
		File: &multipart.FileHeader{
			Filename: "image.jpg",
			Header: make(textproto.MIMEHeader),
			Size: 10,
		},
	}

	mockMenu := domain.Menu {
		Id: mockMenuReq.Id,
		Name: mockMenuReq.Name,
		Price: mockMenuReq.Price,
		Stock: mockMenuReq.Stock,
		Url: "http://localhost:3487/menu/image.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T){
		tempMockMenuReq := &mockMenuReq
		mockMenuRepoPostgres.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(domain.Menu{}, nil).Once()
		mockMenuRepoMinio.On("Upload", mock.AnythingOfType("*multipart.FileHeader"),mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()
		mockMenuRepoPostgres.On("Add", mock.Anything, mock.AnythingOfType("*domain.Menu")).Return(mockMenu.Id, nil).Once()
		mockMenuRepoRedis.On("Delete").Return(nil).Once()

		u := usecase.NewMenuUC(mockMenuRepoPostgres, mockMenuRepoRedis, mockMenuRepoMinio)

		_, err := u.Add(context.TODO(), tempMockMenuReq)

		assert.NoError(t, err)
		mockMenuRepoPostgres.AssertExpectations(t)
	})

	t.Run("failed , reason: 'menu already exist' ", func(t *testing.T){
		tempMockMenuReq := &mockMenuReq
		mockMenuRepoPostgres.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(mockMenu, nil).Once()
		mockMenuRepoMinio.On("Upload", mock.AnythingOfType("*multipart.FileHeader"),mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Maybe()
		mockMenuRepoPostgres.On("Add", mock.Anything, mock.AnythingOfType("*domain.Menu")).Return(mockMenu.Id, nil).Maybe()
		mockMenuRepoRedis.On("Delete").Return(nil).Maybe()

		u := usecase.NewMenuUC(mockMenuRepoPostgres, mockMenuRepoRedis, mockMenuRepoMinio)

		_, err := u.Add(context.TODO(), tempMockMenuReq)

		assert.Error(t, err)
		mockMenuRepoPostgres.AssertExpectations(t)
	})
}
