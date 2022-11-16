package validation

import (
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/models/web"
	validation "github.com/go-ozzo/ozzo-validation"
)

func MenuPayloadValidator(request web.MenuReq){
	err := validation.ValidateStruct(&request, 
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.Stock, validation.Required),
	)

	if err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}
}