package usecase

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/menu/validation"
	"github.com/gofiber/fiber/v2"
)

// menerima repository dan disimpan ke struct menuUseCase
func NewMenuUC(menuRepositoryPostgres domain.MenuRepository, menuRepositoryRedis domain.MenuRepositoryRedis, menuRepositoryMinio domain.MinioRepository) domain.MenuUseCase {
	return &menuUseCaseImpl{
		MenuRepositoryPostgres: menuRepositoryPostgres,
		MenuRepositoryRedis: menuRepositoryRedis,
		MenuRepositoryMinio: menuRepositoryMinio,
	} 
}


// tersimpan repository
type menuUseCaseImpl struct {
	MenuRepositoryPostgres domain.MenuRepository
	MenuRepositoryRedis domain.MenuRepositoryRedis
	MenuRepositoryMinio domain.MinioRepository
}

var bucketName = "menu"
var timestamp = time.Now().Unix()


func (useCase *menuUseCaseImpl) Add(ctx context.Context, request *domain.MenuReq)(response domain.MenuRes, err error) {
	
	if errValidation := validation.MenuPayloadValidator(request); errValidation != nil {
		return response, fiber.NewError(fiber.ErrBadRequest.Code, errValidation.Error())
	}

	objectName :=  strconv.FormatInt(timestamp, 16) +"-"+ request.File.Filename
	
	// Upload file menggunakan Minio
	if err = useCase.MenuRepositoryMinio.Upload(request.File, bucketName, objectName); err != nil {
		return response, err
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
	var getName domain.Menu
	if getName, err = useCase.MenuRepositoryPostgres.GetByName(ctx, menu.Name); err != nil {
		return response, err
	}

	if (getName != domain.Menu{}) {
		return response, fiber.NewError(fiber.ErrBadRequest.Code, "menu is already in use")
	}

	// disimpan ke database
	if _, err = useCase.MenuRepositoryPostgres.Add(ctx, &menu); err != nil {
		return response, err
	}

	if err = useCase.MenuRepositoryRedis.Delete(); err != nil {
		return response, err
	}

	response = domain.MenuRes {
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		Url: menu.Url,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return response, nil
}

func (useCase *menuUseCaseImpl) Delete(ctx context.Context, id string) (response bool,err error){

	var getById domain.Menu
	if getById, err = useCase.MenuRepositoryPostgres.GetById(ctx, id); err != nil {
		return false, err
	}

	if (getById == domain.Menu{}) {
		return false, fiber.NewError(fiber.StatusNotFound, "menu not found")
	}

	if err = useCase.MenuRepositoryPostgres.Delete(ctx, id); err != nil {
		return false, err
	}

	if err = useCase.MenuRepositoryRedis.Delete(); err != nil {
		return false, err
	}

	return true, nil
}

func (useCase *menuUseCaseImpl) GetList(ctx context.Context) (response []domain.MenuRes, err error) {
	
	resultCache, errCache := useCase.MenuRepositoryRedis.Get()

	if errCache != nil {

		var menu []domain.Menu
		if menu, err = useCase.MenuRepositoryPostgres.GetList(ctx); err != nil {
			return nil, err
		}
	
		for _, v := range menu {
			response = append(response, domain.MenuRes{
				Id: v.Id,
				Name: v.Name,
				Price: v.Price,
				Stock: v.Stock,
				Url: v.Url,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}

		_, err = useCase.MenuRepositoryRedis.Set(response)

		if err != nil {
			return nil, err
		}
		
		return response, nil
	}

	
	for _, v := range resultCache {
		response = append(response, domain.MenuRes{
			Id: v.Id,
			Name: v.Name,
			Price: v.Price,
			Stock: v.Stock,
			Url: v.Url,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return response, nil
	
}

func (useCase *menuUseCaseImpl) GetById(ctx context.Context, id string) (response domain.MenuRes, err error) {
	
	var menu domain.Menu
	if menu, err = useCase.MenuRepositoryPostgres.GetById(ctx, id); err != nil {
		return response, err
	}
	
	if (menu == domain.Menu{}) {
		return response, fiber.NewError(fiber.ErrBadRequest.Code, "menu not found")
	}

	response = domain.MenuRes{
		Id: menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
		Url: menu.Url,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	} 
	
	return response, nil
}

func (useCase *menuUseCaseImpl) Update(ctx context.Context, id string, request *domain.MenuReq)(response domain.MenuRes, err error) {

	if errValidation := validation.MenuPayloadValidator(request); errValidation != nil {
		return response, fiber.NewError(fiber.ErrBadRequest.Code, errValidation.Error())
	}

	var menu domain.Menu
	if menu, err = useCase.MenuRepositoryPostgres.GetById(ctx , id); err != nil {
		return response, err
	}

	if(menu == domain.Menu{}) {
		return response, fiber.NewError(fiber.ErrNotFound.Code, "menu not found")
	}

	var urlImage string

	if request.File == nil {
		urlImage = menu.Url
	}else{
		// Upload file menggunakan Minio
		objectName :=  strconv.FormatInt(timestamp, 16) +"-"+ request.File.Filename 

		if err = useCase.MenuRepositoryMinio.Upload(request.File, bucketName, objectName); err != nil {
			return response, err
		}

		urlImage = os.Getenv("MINIO_URL_FILE")+"/"+bucketName+"/"+objectName
	}

	// memindahkan dari request model ke entity/domain Menu
	menuReq := domain.Menu {
		Id: request.Id,
		Name: strings.ToUpper(request.Name),
		Url: urlImage,
		Price: request.Price,
		Stock: request.Stock,
	}

	if _, err = useCase.MenuRepositoryPostgres.Update(ctx, id, &menuReq); err != nil {
		return response, err
	}

	if err = useCase.MenuRepositoryRedis.Delete(); err != nil {
		return response, err
	}
	
	response = domain.MenuRes{
		Id: id,
		Name: menuReq.Name,
		Price: menuReq.Price,
		Stock: menuReq.Stock,
		Url: menuReq.Url,
		CreatedAt: menuReq.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return response, nil
}