package menuusecase

import (
	"strings"

	"github.com/bimaagung/cafe-reservation/exception"
	repository "github.com/bimaagung/cafe-reservation/repository/postgres/menu"

	"github.com/bimaagung/cafe-reservation/models/domain"
	"github.com/bimaagung/cafe-reservation/models/web"
	"github.com/bimaagung/cafe-reservation/validation"
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

func (useCase *menuUseCaseImpl) Add(request web.MenuReq, urlFile string)(response web.MenuRes) {
	
	validation.MenuPayloadValidator(request)

	// memindahkan dari request model ke entity/domain Menu
	menu := domain.Menu {
		Id: request.Id,
		Name: strings.ToUpper(request.Name),
		Url: urlFile,
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

	response = web.MenuRes {
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

func (useCase *menuUseCaseImpl) GetList() (response []web.MenuRes){
	var menus []web.MenuRes

	menu := useCase.MenuRepository.GetList()
	
	for _, v := range menu {
		  menus = append(menus, web.MenuRes{
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

func (useCase *menuUseCaseImpl) GetById(id string) (response web.MenuRes) {
	
	menu := useCase.MenuRepository.GetById(id)
	
	if (menu == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "menu not found",
		})
	}

	response = web.MenuRes{
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	} 
	
	return response
}

func (useCase *menuUseCaseImpl) Update(id string, request web.MenuReq)(response web.MenuRes) {

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
	
	response = web.MenuRes{
		Id: id,
		Name: menuReq.Name,
		Price: menuReq.Price,
		Stock: menuReq.Stock,
		CreatedAt: menuReq.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return response
}