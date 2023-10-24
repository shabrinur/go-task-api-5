package util

import (
	"errors"
	"idstar-idp/rest-api/app/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateUpdateId(id uint) error {
	if id <= 0 {
		return errors.New("field 'id' is required for update")
	}
	return nil
}

func SetErrorResponse(ctx *gin.Context, err error, code int) {
	ctx.AbortWithStatusJSON(code, response.ApiResponse{
		Code:    code,
		Message: err.Error(),
	})
}

func SetSuccessResponseNoData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
	})
}

func SetSuccessResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, response.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}
