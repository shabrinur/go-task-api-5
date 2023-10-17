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

func (ctrl *EmployeeTrainingController) CreateEmployeeTraining(ctx *gin.Context) {
	req := request.EmployeeTrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(false)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.ParseTanggal()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.CreateEmployeeTraining(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeTrainingController) UpdateEmployeeTraining(ctx *gin.Context) {
	req := request.EmployeeTrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(true)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.ParseTanggal()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.UpdateEmployeeTraining(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeTrainingController) GetEmployeeTrainingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetEmployeeTrainingById(id)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeTrainingController) GetEmployeeTrainingList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	validFields := []string{"id", "tanggal", "id_karyawan", "id_training", "created_date", "updated_date"}
	err = req.Validate(validFields)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetEmployeeTrainingList(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeTrainingController) DeleteEmployeeTraining(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = ctrl.svc.DeleteEmployeeTraining(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponseNoData(ctx)
}
