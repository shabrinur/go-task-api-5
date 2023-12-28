package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	service "idstar-idp/rest-api/app/service/emptraining"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	svc *service.EmployeeService
}

func NewEmployeeController(svc *service.EmployeeService) *EmployeeController {
	return &EmployeeController{svc}
}

// CreateEmployee godoc
//
//	@Summary	Create Karyawan
//	@Id			CreateEmployee
//	@Tags		karyawan
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		request	body		request.EmployeeRequest	true	"Create Karyawan Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/idstar/karyawan/save [post]
func (ctrl *EmployeeController) CreateEmployee(ctx *gin.Context) {
	req := request.EmployeeRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.CreateEmployee(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// UpdateEmployee godoc
//
//	@Summary	Update Karyawan
//	@Id			UpdateEmployee
//	@Tags		karyawan
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		request	body		request.EmployeeRequest	true	"Update Karyawan Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	404		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/idstar/karyawan/update [put]
func (ctrl *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	req := request.EmployeeRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.UpdateEmployee(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetEmployeeById godoc
//
//	@Summary	Get Karyawan By Id
//	@Id			GetEmployeeById
//	@Tags		karyawan
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id	path		int	true	"Karyawan ID"
//	@Response	200	{object}	response.ApiResponse
//	@Response	400	{object}	response.ApiResponse
//	@Response	404	{object}	response.ApiResponse
//	@Response	500	{object}	response.ApiResponse
//	@Router		/v1/idstar/karyawan/{id} [get]
func (ctrl *EmployeeController) GetEmployeeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetEmployeeById(id)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetEmployeeList godoc
//
//	@Summary	Get Karyawan List
//	@Id			GetEmployeeList
//	@Tags		karyawan
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		page		query		int		false	"Page"
//	@Param		size		query		int		false	"Size"
//	@Param		field		query		string	false	"Field"
//	@Param		direction	query		string	false	"Direction"
//	@Response	200			{object}	response.PaginationData
//	@Response	400			{object}	response.ApiResponse
//	@Response	500			{object}	response.ApiResponse
//	@Router		/v1/idstar/karyawan/list [get]
func (ctrl *EmployeeController) GetEmployeeList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetEmployeeList(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// DeleteEmployee godoc
//
//	@Summary	Delete Karyawan
//	@Id			DeleteEmployee
//	@Tags		karyawan
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		request	body		request.IdRequest	true	"Delete Karyawan Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/idstar/karyawan/delete [delete]
func (ctrl *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	code, err := ctrl.svc.DeleteEmployee(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponseNoData(ctx)
}
