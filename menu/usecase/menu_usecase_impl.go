package usecase

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	repoitorypostgres "github.com/bimaagung/cafe-reservation/menu/repository/postgres"
	repoitoryredis "github.com/bimaagung/cafe-reservation/menu/repository/redis"
	"github.com/bimaagung/cafe-reservation/menu/validation"
	minioUpload "github.com/bimaagung/cafe-reservation/pkg/minio"
	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/gofiber/fiber/v2"
)

// menerima repository dan disimpan ke struct menuUseCase
func NewMenuUC(menuRepositoryPostgres *repoitorypostgres.MenuRepository,menuRepositoryRedis *repoitoryredis.MenuRepositoryRedis) MenuUseCase {
	return &menuUseCaseImpl{
		MenuRepositoryPostgres: *menuRepositoryPostgres,
		MenuRepositoryRedis: *menuRepositoryRedis,
	} 
}


// tersimpan repository
type menuUseCaseImpl struct {
	MenuRepositoryPostgres repoitorypostgres.MenuRepository
	MenuRepositoryRedis repoitoryredis.MenuRepositoryRedis
}

// TODO : url response is empty
func (useCase *menuUseCaseImpl) Add(ctx *fiber.Ctx, request domain.MenuReq)(response domain.MenuRes) {
	
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
	getName := useCase.MenuRepositoryPostgres.GetByName(ctx, menu.Name)

	if (getName != domain.Menu{}) {
		panic(exception.NotFoundError{
			Message: "menu is already in use",
		})
	}

	// disimpan ke database
	useCase.MenuRepositoryPostgres.Add(ctx, menu)

	errCache := useCase.MenuRepositoryRedis.Delete()
	exception.Error(errCache)

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

func (useCase *menuUseCaseImpl) Delete(ctx *fiber.Ctx, id string) bool{

	getById := useCase.MenuRepositoryPostgres.GetById(ctx, id)
	
	if (getById == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "menu not found",
		})
	}

	useCase.MenuRepositoryPostgres.Delete(ctx, id)

	errCache := useCase.MenuRepositoryRedis.Delete()
	exception.Error(errCache)

	response := true
	return response
}

func (useCase *menuUseCaseImpl) GetList(ctx *fiber.Ctx) (response []domain.MenuRes){
	var menus []domain.MenuRes
	
	resultCache, errCache := useCase.MenuRepositoryRedis.Get()

	if errCache != nil {
		menu := useCase.MenuRepositoryPostgres.GetList(ctx)
	
		for _, v := range menu {
			menus = append(menus, domain.MenuRes{
				Id: v.Id,
				Name: v.Name,
				Price: v.Price,
				Stock: v.Stock,
				Url: v.Url,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}

		_, errCache := useCase.MenuRepositoryRedis.Set(menus)

		if errCache != nil {
			panic(exception.ClientError{
				Message: errCache.Error(),
			})
		}
		
		return menus
	}

	
	for _, v := range resultCache {
		menus = append(menus, domain.MenuRes{
			Id: v.Id,
			Name: v.Name,
			Price: v.Price,
			Stock: v.Stock,
			Url: v.Url,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return menus
	
}


func (useCase *menuUseCaseImpl) GetById(ctx *fiber.Ctx, id string) (response domain.MenuRes) {
	
	menu := useCase.MenuRepositoryPostgres.GetById(ctx, id)
	
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

// TODO : fixing update request
func (useCase *menuUseCaseImpl) Update(ctx *fiber.Ctx, id string, request domain.MenuReq)(response domain.MenuRes) {

	menuReq :=  domain.Menu{
		Name: strings.ToUpper(request.Name),
		Price: request.Price,
		Stock: request.Stock,
	}

	menu := useCase.MenuRepositoryPostgres.GetById(ctx , id)

	if(menu == domain.Menu{}) {
		panic(exception.ClientError{
			Message: "menu not found",
		})
	}


	useCase.MenuRepositoryPostgres.Update(ctx, id, menuReq)

	errCache := useCase.MenuRepositoryRedis.Delete()
	exception.Error(errCache)

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