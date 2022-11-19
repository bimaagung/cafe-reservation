package usecase

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/menu/repository"
	"github.com/bimaagung/cafe-reservation/menu/validation"
	minioUpload "github.com/bimaagung/cafe-reservation/pkg/minio"

	"github.com/bimaagung/cafe-reservation/utils/exception"
)

// menerima repository dan disimpan ke struct menuUseCase
func NewMenuUC(menuRepository *repository.MenuRepository) MenuUseCase {
	return &menuUseCaseImpl{
		MenuRepository: *menuRepository,
	} 
}


// tersimpan repository
type menuUseCaseImpl struct {
	MenuRepository repository.MenuRepository
}

func (useCase *menuUseCaseImpl) Add(request domain.MenuReq)(response domain.MenuRes) {
	
	validation.MenuPayloadValidator(request)

	bucketName := "menu"
	timestamp := time.Now().Unix()
	objectName :=  strconv.FormatInt(timestamp, 16) +"-"+ request.File.Filename
	
	// Upload file menggunakan Minio
	errUpload := minioUpload.UploadFile(request.File, bucketName, objectName)

	if errUpload != nil {
		panic(exception.ClientError{
			Message: errUpload.Error(),
		})
	}

	// memindahkan dari request model ke entity/domain Menu
	menu := domain.Menu {
		Id: request.Id,
		Name: strings.ToUpper(request.Name),
		Url: os.Getenv("MINIO_URL_FILE")+"/"+bucketName+"/"+objectName,
		Price: request.Price,
		Stock: request.Stock,
	}

	// validasi menu apabila menu sudah ada
	getName := useCase.MenuRepository.GetByName(menu.Name)

	if (getName != domain.Menu{}) {
		panic(exception.NotFoundError{
			Message: "menu is already in use",
		})
	}

	// disimpan ke database
	useCase.MenuRepository.Add(menu)

	response = domain.MenuRes {
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		Url: menu.Url,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return response
}

func (useCase *menuUseCaseImpl) Delete(id string) bool{

	getById := useCase.MenuRepository.GetById(id)
	
	if (getById == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "menu not found",
		})
	}

	useCase.MenuRepository.Delete(id)

	response := true
	return response
}

func (useCase *menuUseCaseImpl) GetList() (response []domain.MenuRes){
	var menus []domain.MenuRes

	menu := useCase.MenuRepository.GetList()
	
	for _, v := range menu {
		  menus = append(menus, domain.MenuRes{
			Id: v.Id,
			Name: v.Name,
			Price: v.Price,
			Stock: v.Stock,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		  })
	}
	

	return menus
}

func (useCase *menuUseCaseImpl) GetById(id string) (response domain.MenuRes) {
	
	menu := useCase.MenuRepository.GetById(id)
	
	if (menu == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "menu not found",
		})
	}

	response = domain.MenuRes{
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	} 
	
	return response
}

func (useCase *menuUseCaseImpl) Update(id string, request domain.MenuReq)(response domain.MenuRes) {

	menuReq :=  domain.Menu{
		Name: strings.ToUpper(request.Name),
		Price: request.Price,
		Stock: request.Stock,
	}

	menu := useCase.MenuRepository.GetById(id)

	if(menu == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "menu not found",
		})
	}


	useCase.MenuRepository.Update(id, menuReq)

	validation.MenuPayloadValidator(request)
	
	response = domain.MenuRes{
		Id: id,
		Name: menuReq.Name,
		Price: menuReq.Price,
		Stock: menuReq.Stock,
		CreatedAt: menuReq.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return response
}