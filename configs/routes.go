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

	/*-----------------------------------------------------------------------------------------
		Local Government CRU
	-------------------------------------------------------------------------------------------*/
	route.POST("/localgovernment/create", func(ctx *gin.Context) {
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
	route.GET("/localgovernment", func(ctx *gin.Context) {

		//default http status code ==200
		code := http.StatusOK
		pagination := helpers.GeneratePaginationRequest(ctx)
		response := services.FindAllLocalGovernment(*localGovernmentRepository, ctx, pagination)
		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)

	})

	route.GET("/localgovernment/show/state/:id", func(ctx *gin.Context) {
		//default http status code ==200
		code := http.StatusOK

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {

			return
		}
		pagination := helpers.GeneratePaginationRequest(ctx)

		response := services.FindLocalGovernmentByState(id, *localGovernmentRepository, ctx, pagination)
		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})
	route.PUT("/localgovernment/update/:id", func(ctx *gin.Context) {

		var localgovernment models.LocalGovernment

		//validate json
		err := ctx.ShouldBindJSON(&localgovernment)

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
		response := services.UpdateLocalGovernmentById(id, &localgovernment, *localGovernmentRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})
	/*--------------------------------------------------------------------------------------
	  WARD CRU
	  ----------------------------------------------------------------------------------------*/

	route.POST("/ward/create", func(ctx *gin.Context) {
		var ward models.Ward

		//validate json
		err := ctx.ShouldBindJSON(&ward)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		//default http status code ==200
		code := http.StatusOK

		//save state
		response := services.CreateWard(&ward, *wardRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})

	route.GET("/ward", func(ctx *gin.Context) {

		//default http status code ==200
		code := http.StatusOK
		pagination := helpers.GeneratePaginationRequest(ctx)
		response := services.FindAllWard(*wardRepository, ctx, pagination)
		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)

	})

	route.GET("/ward/show/localgovernment/:id", func(ctx *gin.Context) {
		//default http status code ==200
		code := http.StatusOK

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {

			return
		}
		pagination := helpers.GeneratePaginationRequest(ctx)

		response := services.FindWardByLocalGovernment(id, *wardRepository, ctx, pagination)
		if !response.Success {
			code = http.StatusBadRequest
		}
		ctx.JSON(code, response)
	})
	/*-------------------------------------------------------------------------------------
		Return Endpoints
	---------------------------------------------------------------------------------------*/
	return route
}
