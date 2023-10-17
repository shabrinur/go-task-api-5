package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/service"
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

func (ctrl *EmployeeController) CreateEmployee(ctx *gin.Context) {
	req := request.EmployeeRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(false)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.ParseDob()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.CreateEmployee(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	req := request.EmployeeRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(true)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.ParseDob()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.UpdateEmployee(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeController) GetEmployeeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetEmployeeById(id)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeController) GetEmployeeList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	validFields := []string{"id", "nama", "status", "dob", "detail_karyawan", "created_date", "updated_date"}
	err = req.Validate(validFields)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetEmployeeList(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = ctrl.svc.DeleteEmployee(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponseNoData(ctx)
}
