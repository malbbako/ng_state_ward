package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/malbbako/ng_state_ward/dtos"
	"github.com/malbbako/ng_state_ward/langs"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.Validation

	//get validation errors
	ValidationErrors := err.(validator.ValidationErrors)
	for _, value := range ValidationErrors {
		//get field & rule
		field, rule := value.Field(), value.Tag()

		//create validation object
		validation := dtos.Validation{Field: field, Message: langs.GenerateValidMessage(field, rule)}

		//add validation object to validations list
		validations = append(validations, validation)

	}
	//set response validations
	response.Validations = validations

	return response
}
