package usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/user/repository"
	tokenmanager "github.com/bimaagung/cafe-reservation/utils/token_manager"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func NewUserUC(userRepository repository.UserRepository) UserUseCase {
	return &userUseCaseImpl {
		UserRepository: userRepository,
	}
}

type userUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func (useCase *userUseCaseImpl) Create(ctx context.Context, request *userdomain.UserReq)(response userdomain.UserRes, err error){

	// Check match password
	if request.Password != request.RetypePassword {
		return response, fiber.NewError(fiber.ErrBadRequest.Code, "password and retype password not match")
	}

	// Hash Password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return response, err
	}
	
	user := userdomain.User{
		Id: request.Id,
		Name: request.Name,
		Username: request.Username,
		Password: string(hashPassword),
	}

	// Check usernam already exists
	if userByUsername, err := useCase.UserRepository.GetByUsername(ctx, user.Username); userByUsername.Id != "" {
		if err != nil {
			return response, err	
		}
		return response, fiber.NewError(fiber.ErrBadRequest.Code, "user already exists")
	}
	
	if _, err = useCase.UserRepository.Create(ctx, &user); err != nil {
		return response, err
	}

	expTime , err := strconv.Atoi(os.Getenv("EXPIRED_TOKEN"))
	if err != nil {
		return response, err
	}

	//Generate Token
	claims := jwt.MapClaims{
		"id": request.Id,
		"name": request.Name,
		"username": request.Username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(expTime) * time.Hour).Unix(),
		"iss": os.Getenv("APP_NAME"),
	}

	token, err := tokenmanager.GenerateToken(claims) 

	if err != nil {
		return response, err
	}

	response = userdomain.UserRes{
		Id: request.Id,
		Name: request.Name,
		Username: request.Username,
		Token: token,
	}

	return response, nil
}