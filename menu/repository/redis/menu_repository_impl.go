package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/go-redis/redis/v8"
)

func NewRepositoryRedis(redisDB *redis.Client) MenuRepositoryRedis {
	return &redisMenuRepository{
		redisDB: redisDB,
	} 
}

type redisMenuRepository struct {
	redisDB *redis.Client
}

var ctx = context.Background()

func (repository *redisMenuRepository) Set(menu interface{}) (string, error) {  

	valEncode, errEncode := json.Marshal(menu) 

	if errEncode != nil {
		return "", errEncode
	} 

	val, err := repository.redisDB.Set(ctx, "menu", valEncode, 0).Result()
	
	if err != nil {
		return "", err
	}

	return val, nil
}

func (repository *redisMenuRepository) Delete() error {  
	err := repository.redisDB.Del(ctx, "menu").Err()
	return err
}


func (repository *redisMenuRepository) Get()([]domain.Menu, error) {  
	var menus []domain.Menu

	val, errResult := repository.redisDB.Get(ctx, "menu").Result()

	if errResult == redis.Nil {
		return nil, errors.New("menu not found")
	} else if errResult != nil {
		return nil, errResult
	}else{
		errDecode := json.Unmarshal([]byte(val), &menus)

		if errDecode != nil {
			return nil, errResult
		}

		return menus, nil
	}

}

