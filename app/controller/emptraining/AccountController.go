package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	service "idstar-idp/rest-api/app/service/emptraining"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	svc *service.AccountService
}

func NewAccountController(svc *service.AccountService) *AccountController {
	return &AccountController{svc}
}

// CreateAccount godoc
//
//	@Summary	Create Rekening
//	@Id			CreateAccount
//	@Tags		rekening
//	@Accept		json
//	@Produce	json
//	@Param		request	body		request.AccountRequest	true	"Create Rekening Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/idstar/rekening/save [post]
func (ctrl *AccountController) CreateAccount(ctx *gin.Context) {
	req := request.AccountRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.CreateAccount(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// UpdateAccount godoc
//
//	@Summary	Update Rekening
//	@Id			UpdateAccount
//	@Tags		rekening
//	@Accept		json
//	@Produce	json
//	@Param		request	body		request.AccountRequest	true	"Update Rekening Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	404		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/idstar/rekening/update [put]
func (ctrl *AccountController) UpdateAccount(ctx *gin.Context) {
	req := request.AccountRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.UpdateAccount(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetAccountById godoc
//
//	@Summary	Get Rekening By Id
//	@Id			GetAccountById
//	@Tags		rekening
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Rekening ID"
//	@Response	200	{object}	response.ApiResponse
//	@Response	400	{object}	response.ApiResponse
//	@Response	404	{object}	response.ApiResponse
//	@Response	500	{object}	response.ApiResponse
//	@Router		/v1/idstar/rekening/{id} [get]
func (ctrl *AccountController) GetAccountById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetAccountById(id)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetAccountList godoc
//
//	@Summary	Get Rekening List
//	@Id			GetAccountList
//	@Tags		rekening
//	@Accept		json
//	@Produce	json
//	@Param		page		query		int		false	"Page"
//	@Param		size		query		int		false	"Size"
//	@Param		field		query		string	false	"Field"
//	@Param		direction	query		string	false	"Direction"
//	@Response	200			{object}	response.PaginationData
//	@Response	400			{object}	response.ApiResponse
//	@Response	500			{object}	response.ApiResponse
//	@Router		/v1/idstar/rekening/list [get]
func (ctrl *AccountController) GetAccountList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetAccountList(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// DeleteAccount godoc
//
//	@Summary	Delete Rekening
//	@Id			DeleteAccount
//	@Tags		rekening
//	@Accept		json
//	@Produce	json
//	@Param		request	body		request.IdRequest	true	"Delete Rekening Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/idstar/rekening/delete [delete]
func (ctrl *AccountController) DeleteAccount(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	code, err := ctrl.svc.DeleteAccount(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponseNoData(ctx)
}
