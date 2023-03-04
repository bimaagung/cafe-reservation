package validation

import (
	"github.com/bimaagung/cafe-reservation/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

func MenuPayloadValidator(request *domain.MenuReq) error {
	err := validation.ValidateStruct(&request, 
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.Stock, validation.Required),
	)

	if err != err {
		return err
	}

	return nil
}