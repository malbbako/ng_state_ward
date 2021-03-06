package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/malbbako/ng_state_ward/dtos"
	"github.com/malbbako/ng_state_ward/models"
	"github.com/malbbako/ng_state_ward/repositories"
)

func CreateState(state *models.State, repository repositories.StateRepository) dtos.Response {
	// uuidResult, err := uuid.NewRandom()

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// state.ID = uuidResult.String()
	operationResult := repository.Save(state)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	data := operationResult.Result.(*models.State)
	return dtos.Response{Success: true, Data: data}
}

func FindAllStates(repository repositories.StateRepository, ctx *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.FindAll(pagination)
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
func FindOneById(id int, repository repositories.StateRepository) dtos.Response {

	operationResult := repository.FindOneById(id)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.State)
	return dtos.Response{Success: true, Data: data}
}
func UpdateStateById(id int, state *models.State, repository repositories.StateRepository) dtos.Response {

	existingStateResponse := FindOneById(id, repository)
	if !existingStateResponse.Success {
		return existingStateResponse
	}
	existingState := existingStateResponse.Data.(*models.State)
	existingState.Name = state.Name
	existingState.Abbr = state.Abbr

	operationResult := repository.Save(existingState)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}

}
