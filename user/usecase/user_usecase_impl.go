package usecase

import (
	"os"
	"strconv"
	"time"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/user/repository"
	"github.com/bimaagung/cafe-reservation/utils/exception"
	tokenmanager "github.com/bimaagung/cafe-reservation/utils/token_manager"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func NewUserUC(userRepository *repository.UserRepository) UserUseCase {
	return &userUseCaseImpl {
		UserRepository: *userRepository,
	}
}

type userUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func (useCase *userUseCaseImpl) Create(ctx *fiber.Ctx, request userdomain.UserReq)(response userdomain.UserRes){

	// Check match password
	if request.Password != request.RetypePassword {
		panic(exception.ClientError{
			Message: "password and retype password not match",
		})
	}
	// Hash Password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		exception.Error(err.Error())
	}
	
	user := userdomain.User{
		Id: request.Id,
		Name: request.Name,
		Username: request.Username,
		Password: string(hashPassword),
	}

	// Create User
	useCase.UserRepository.Create(ctx, user)

	expTime , _ := strconv.Atoi(os.Getenv("EXPIRED_TOKEN"))

	//Generate Token
	claims := jwt.MapClaims{
		"id": request.Id,
		"name": request.Name,
		"username": request.Username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(expTime) * time.Hour).Unix(),
		"iss": os.Getenv("APP_NAME"),
	}

	token := tokenmanager.GenerateToken(claims)

	response = userdomain.UserRes{
		Id: request.Id,
		Name: request.Name,
		Username: request.Username,
		 Token: token,
	}

	return response
}