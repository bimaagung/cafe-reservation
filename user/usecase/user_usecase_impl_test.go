package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/user/mocks"
	"github.com/bimaagung/cafe-reservation/user/usecase"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const patchEnv = "../../.env"

func TestUserUC_Create(t *testing.T) {
	godotenv.Load(patchEnv)
	
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.UserReq{
		Id:        "97391bbb-a48f-48e2-a166-db669e6377fc",
		Name:      "Jo Sauer",
		Username:  "Aracely_Simonis",
		Password:  "12345678",
		RetypePassword: "12345678",
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := &mockUser
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, nil).Once()
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return("97391bbb-a48f-48e2-a166-db669e6377fc", nil).Once()

		u := usecase.NewUserUC(mockUserRepo)

		_, err := u.Create(context.TODO(), tempMockUser)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("user exists", func(t *testing.T) {
		resultUser := domain.User{
			Id:        "97391bbb-a48f-48e2-a166-db669e6377fc",
			Name:      "Jo Sauer",
			Username:  "Aracely_Simonis",
			Password:  "12345678",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		tempMockUser := mockUser
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(resultUser, nil).Once()
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.User")).Return("97391bbb-a48f-48e2-a166-db669e6377fc", nil).Maybe()

		u := usecase.NewUserUC(mockUserRepo)

		_, err := u.Create(context.TODO(), &tempMockUser)
		
		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
