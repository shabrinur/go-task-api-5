package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/service"
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

func (ctrl *AccountController) CreateAccount(ctx *gin.Context) {
	req := request.AccountRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(false)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.CreateAccount(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *AccountController) UpdateAccount(ctx *gin.Context) {
	req := request.AccountRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(true)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.UpdateAccount(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *AccountController) GetAccountById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetAccountById(id)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *AccountController) GetAccountList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	validFields := []string{"id", "nama", "jenis", "rekening", "id_karyawan", "created_date", "updated_date"}
	err = req.Validate(validFields)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetAccountList(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *AccountController) DeleteAccount(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = ctrl.svc.DeleteAccount(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponseNoData(ctx)
}
