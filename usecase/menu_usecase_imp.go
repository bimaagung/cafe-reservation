package usecase

import (
	"time"

	"github.com/bimaagung/cafe-reservation/entity"
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/model"
	"github.com/bimaagung/cafe-reservation/repository"
)

func NewMenuUC(menuRepository *repository.MenuRepository) MenuUseCase {
	return &menuUseCaseImpl{
		MenuRepository: *menuRepository,
	} 
}

type menuUseCaseImpl struct {
	MenuRepository repository.MenuRepository
}

func (useCase *menuUseCaseImpl) Add(request model.Menu)(response model.Menu) {

	menu := entity.Menu {
		Id: request.Id,
		Name: request.Name,
		Price: request.Price,
		Stock: request.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	getName := useCase.MenuRepository.GetByName(menu.Name)

	if (getName != entity.Menu{}) {
		panic(exception.NotFoundError{
			Message: "category is already in use",
		})
	}

	useCase.MenuRepository.Add(menu)

	response = model.Menu {
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		CreatedAt: menu.CreatedAt.String(),
		UpdatedAt: menu.UpdatedAt.String(),
	}

	return response
}

func (useCase *menuUseCaseImpl) Delete(id string) bool{

	getById := useCase.MenuRepository.GetById(id)
	
	if (getById == entity.Menu{}) {
		panic(exception.ClientError{
			Message: "category not found",
		})
	}

	useCase.MenuRepository.Delete(id)

	response := true
	return response
}

func (useCase *menuUseCaseImpl) GetList() (response []model.Menu){
	var menus []model.Menu

	menu := useCase.MenuRepository.GetList()
	
	for _, v := range menu {
		  menus = append(menus, model.Menu{
			Id: v.Id,
			Name: v.Name,
			Price: v.Price,
			Stock: v.Stock,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
		  })
	}
	

	return menus
}