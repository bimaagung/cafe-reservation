package usecase

import (
	"os"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/user/repository"
	"github.com/bimaagung/cafe-reservation/utils/exception"
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

func (useCase *userUseCaseImpl) Create(request userdomain.UserReq)(response userdomain.UserRes){

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
	useCase.UserRepository.Create(user)

	//Generate Token
	claims := jwt.MapClaims{
		"id": request.Id,
		"name": request.Name,
		"username": request.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		exception.Error(err.Error())
	}

	response = userdomain.UserRes{
		Id: request.Id,
		Name: request.Name,
		Username: request.Username,
		 Token: t,
	}

	return response
}