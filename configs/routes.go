package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malbbako/ng_state_ward/helpers"
	"github.com/malbbako/ng_state_ward/models"
	"github.com/malbbako/ng_state_ward/repositories"
	"github.com/malbbako/ng_state_ward/services"
)

func SetupRoutes(stateRepository *repositories.StateRepository) *gin.Engine {
	route := gin.Default()

	route.POST("/state/create", func(ctx *gin.Context) {
		var state models.State

		//validate json
		err := ctx.ShouldBindJSON(&state)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		//default http status code ==200
		code := http.StatusOK

		//save state
		response := services.CreateState(&state, *stateRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})

	return route
}
