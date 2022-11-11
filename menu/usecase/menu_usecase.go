package usecase

import (
	"time"

	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/exception"
)

func NewMenuUC(menuRepository *domain.MenuRepository) domain.MenuUseCase {
	return &menuUseCaseImpl{
		MenuRepository: *menuRepository,
	} 
}

type menuUseCaseImpl struct {
	MenuRepository domain.MenuRepository
}

func (useCase *menuUseCaseImpl) Add(request domain.Menu)(response domain.Menu) {

	menu := domain.Menu {
		Id: request.Id,
		Name: request.Name,
		Price: request.Price,
		Stock: request.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	getName := useCase.MenuRepository.GetByName(menu.Name)

	if (getName != domain.Menu{}) {
		panic(exception.NotFoundError{
			Message: "category is already in use",
		})
	}

	useCase.MenuRepository.Add(menu)

	response = domain.Menu {
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return response
}

func (useCase *menuUseCaseImpl) Delete(id string) bool{

	getById := useCase.MenuRepository.GetById(id)
	
	if (getById == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "category not found",
		})
	}

	useCase.MenuRepository.Delete(id)

	response := true
	return response
}

func (useCase *menuUseCaseImpl) GetList() (response []domain.Menu){
	var menus []domain.Menu

	menu := useCase.MenuRepository.GetList()
	
	for _, v := range menu {
		  menus = append(menus, domain.Menu{
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