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

func AddMenuUseCase_Test(t *testing.T) {
	godotenv.Load(patchEnv)

	mockMenuRepoPostgres := new(mocks.MenuRepositoryPostgresMock) 
	mockMenuRepoRedis := new(mocks.MenuRepositoryRedisMock) 
	
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
		tempMockMenu := &mockMenu
		mockMenuRepoPostgres.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(tempMockMenu, nil)
		mockMenuRepoPostgres.On("Add", mock.Anything, mock.AnythingOfType("*domain.Menu")).Return(tempMockMenu.Id, nil)
		mockMenuRepoRedis.On("Delete").Return(nil)

		u := usecase.NewMenuUC(mockMenuRepoPostgres, mockMenuRepoRedis)

		_, err := u.Add(context.TODO(), tempMockMenuReq)

		assert.NoError(t, err)
		mockMenuRepoPostgres.AssertExpectations(t)
	})
}