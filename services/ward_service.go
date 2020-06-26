package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/malbbako/ng_state_ward/dtos"
	"github.com/malbbako/ng_state_ward/models"
	"github.com/malbbako/ng_state_ward/repositories"
)

func CreateWard(ward *models.Ward, repository repositories.WardRepository) dtos.Response {
	// uuidResult, err := uuid.NewRandom()

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// state.ID = uuidResult.String()
	operationResult := repository.CreateWard(ward)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	data := operationResult.Result.(*models.Ward)
	return dtos.Response{Success: true, Data: data}
}

func FindAllWard(repository repositories.WardRepository, ctx *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.FindAllWard(pagination)
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
func FindWardByLocalGovernment(id int, repository repositories.WardRepository, ctx *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.FindWardByLocalgovernment(id, pagination)
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
