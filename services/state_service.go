package services

import (
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

func FindAllStates(repository repositories.StateRepository) dtos.Response {
	operationResult := repository.FindAll()
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.States)
	return dtos.Response{Success: true, Data: data}
}
