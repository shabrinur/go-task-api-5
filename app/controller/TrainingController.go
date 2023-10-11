package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/service"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrainingController struct {
	svc *service.TrainingService
}

func NewTrainingController(svc *service.TrainingService) *TrainingController {
	return &TrainingController{svc}
}

func (ctrl *TrainingController) CreateTraining(ctx *gin.Context) {
	req := request.TrainingRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := req.Validate(false); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ctrl.svc.CreateTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) UpdateTraining(ctx *gin.Context) {
	req := request.TrainingRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := req.Validate(true); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ctrl.svc.UpdateTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) GetTrainingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ctrl.svc.GetTrainingById(id)
	if err != nil {
		util.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) GetTrainingList(ctx *gin.Context) {
	req := request.PagingRequest{}

	if err := ctx.Bind(&req); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	validFields := []string{"id", "tema", "pengajar", "created_date", "updated_date"}
	if err := req.Validate(validFields); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ctrl.svc.GetTrainingList(req)
	if err != nil {
		util.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) DeleteTraining(ctx *gin.Context) {
	req := request.IdRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		util.SetErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := ctrl.svc.DeleteTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	util.SetSuccessResponseNoData(ctx)
}
