package configs

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malbbako/ng_state_ward/helpers"
	"github.com/malbbako/ng_state_ward/models"
	"github.com/malbbako/ng_state_ward/repositories"
	"github.com/malbbako/ng_state_ward/services"
)

func SetupRoutes(stateRepository *repositories.StateRepository, localGovernmentRepository *repositories.LocalGovernmentRepository, wardRepository *repositories.WardRepository) *gin.Engine {
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
	route.GET("/state", func(ctx *gin.Context) {

		//default http status code ==200
		code := http.StatusOK
		pagination := helpers.GeneratePaginationRequest(ctx)
		response := services.FindAllStates(*stateRepository, ctx, pagination)
		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)

	})
	route.GET("/state/show/:id", func(ctx *gin.Context) {
		//default http status code ==200
		code := http.StatusOK

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {

			return
		}
		response := services.FindOneById(id, *stateRepository)
		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)

	})

	route.PUT("/state/update/:id", func(ctx *gin.Context) {

		var state models.State

		//validate json
		err := ctx.ShouldBindJSON(&state)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {

			return
		}

		code := http.StatusOK

		//save state
		response := services.UpdateStateById(id, &state, *stateRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})

	route.POST("/localgov/create", func(ctx *gin.Context) {
		var localGovernment models.LocalGovernment

		//validate json
		err := ctx.ShouldBindJSON(&localGovernment)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		//default http status code ==200
		code := http.StatusOK

		//save state
		response := services.CreateLocalGovernment(&localGovernment, *localGovernmentRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})

	return route
}
