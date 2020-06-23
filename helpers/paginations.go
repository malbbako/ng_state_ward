package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malbbako/ng_state_ward/dtos"
)

func GeneratePaginationRequest(ctx *gin.Context) *dtos.Pagination {
	//convert query parameter string to int
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	sort := ctx.DefaultQuery("sort", "created_at desc")
	return &dtos.Pagination{Limit: limit, Page: page, Sort: sort}
}
