package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/malbbako/ng_state_ward/dtos"
	"github.com/malbbako/ng_state_ward/models"
	"github.com/malbbako/ng_state_ward/repositories"
)

func CreateLocalGovernment(localgovernment *models.LocalGovernment, repository repositories.LocalGovernmentRepository) dtos.Response {
	// uuidResult, err := uuid.NewRandom()

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// state.ID = uuidResult.String()
	operationResult := repository.CreateLocalGovernment(localgovernment)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	data := operationResult.Result.(*models.LocalGovernment)
	return dtos.Response{Success: true, Data: data}
}

func FindAllLocalGovernment(repository repositories.LocalGovernmentRepository, ctx *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.FindAllLocalGovernment(pagination)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*dtos.Pagination)
	//get current url
	urlPath := ctx.Request.URL.Path
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort)
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort)

	if data.Page > 0 {
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort)

	}

	if data.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort)

	}

	if data.Page > totalPages {
		data.FirstPage = ""
	}
	return dtos.Response{Success: true, Data: data}
}
func FindLocalGovernmentById(id int, repository repositories.LocalGovernmentRepository) dtos.Response {

	operationResult := repository.FindLocalGovernmentById(id)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.LocalGovernment)
	return dtos.Response{Success: true, Data: data}
}
func UpdateLocalGovernmentById(id int, localGovernment *models.LocalGovernment, repository repositories.LocalGovernmentRepository) dtos.Response {

	existingLocalGovernmentResponse := FindLocalGovernmentById(id, repository)
	if !existingLocalGovernmentResponse.Success {
		return existingLocalGovernmentResponse
	}
	existingLocalGovernment := existingLocalGovernmentResponse.Data.(*models.LocalGovernment)
	existingLocalGovernment.Name = localGovernment.Name
	existingLocalGovernment.Abbr = localGovernment.Abbr

	operationResult := repository.Save(existingLocalGovernment)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}

}
