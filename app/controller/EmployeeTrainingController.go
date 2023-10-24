package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/service"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeTrainingController struct {
	svc *service.EmployeeTrainingService
}

func NewEmployeeTrainingController(svc *service.EmployeeTrainingService) *EmployeeTrainingController {
	return &EmployeeTrainingController{svc}
}

// CreateEmployeeTraining godoc
//
//	@Summary	Create Karyawan Training
//	@Id			CreateEmployeeTraining
//	@Tags		karyawan-training
//	@Accept		json
//	@Produce	json
//	@Param		request	body		request.EmployeeTrainingRequest	true	"Create Karyawan Training Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/karyawan-training/save [post]
func (ctrl *EmployeeTrainingController) CreateEmployeeTraining(ctx *gin.Context) {
	req := request.EmployeeTrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.CreateEmployeeTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// UpdateEmployeeTraining godoc
//
//	@Summary	Update Karyawan Training
//	@Id			UpdateEmployeeTraining
//	@Tags		karyawan-training
//	@Accept		json
//	@Produce	json
//	@Param		request	body		request.EmployeeTrainingRequest	true	"Update Karyawan Training Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	404		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/karyawan-training/update [put]
func (ctrl *EmployeeTrainingController) UpdateEmployeeTraining(ctx *gin.Context) {
	req := request.EmployeeTrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.UpdateEmployeeTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetEmployeeTrainingById godoc
//
//	@Summary	Get Karyawan Training By Id
//	@Id			GetEmployeeTrainingById
//	@Tags		karyawan-training
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Training ID"
//	@Response	200	{object}	response.ApiResponse
//	@Response	400	{object}	response.ApiResponse
//	@Response	404	{object}	response.ApiResponse
//	@Response	500	{object}	response.ApiResponse
//	@Router		/karyawan-training/{id} [get]
func (ctrl *EmployeeTrainingController) GetEmployeeTrainingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetEmployeeTrainingById(id)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetEmployeeTrainingList godoc
//
//	@Summary	Get Karyawan Training List
//	@Id			GetEmployeeTrainingList
//	@Tags		karyawan-training
//	@Accept		json
//	@Produce	json
//	@Param		page		query		int		false	"Page"
//	@Param		size		query		int		false	"Size"
//	@Param		field		query		string	false	"Field"
//	@Param		direction	query		string	false	"Direction"
//	@Response	200			{object}	response.PaginationData
//	@Response	400			{object}	response.ApiResponse
//	@Response	500			{object}	response.ApiResponse
//	@Router		/karyawan-training/list [get]
func (ctrl *EmployeeTrainingController) GetEmployeeTrainingList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetEmployeeTrainingList(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// DeleteEmployeeTraining godoc
//
//	@Summary	Delete Karyawan Training
//	@Id			DeleteEmployeeTraining
//	@Tags		karyawan-training
//	@Accept		json
//	@Produce	json
//	@Param		request	body		request.IdRequest	true	"Delete Karyawan Training Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/karyawan-training/delete [delete]
func (ctrl *EmployeeTrainingController) DeleteEmployeeTraining(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	code, err := ctrl.svc.DeleteEmployeeTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponseNoData(ctx)
}
