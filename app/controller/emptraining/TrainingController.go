package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	service "idstar-idp/rest-api/app/service/emptraining"
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

// CreateTraining godoc
//
//	@Summary		Create Training
//	@Description	To authorize this API, get token from user-login
//	@Id				CreateTraining
//	@Tags			idstar/training
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Param			request	body		request.TrainingRequest	true	"Create Training Request"
//	@Response		200		{object}	response.ApiResponse
//	@Response		400		{object}	response.ApiResponse
//	@Response		500		{object}	response.ApiResponse
//	@Router			/v1/idstar/training/save [post]
func (ctrl *TrainingController) CreateTraining(ctx *gin.Context) {
	req := request.TrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.CreateTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// UpdateTraining godoc
//
//	@Summary		Update Training
//	@Description	To authorize this API, get token from user-login
//	@Id				UpdateTraining
//	@Tags			idstar/training
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Param			request	body		request.TrainingRequest	true	"Update Training Request"
//	@Response		200		{object}	response.ApiResponse
//	@Response		400		{object}	response.ApiResponse
//	@Response		404		{object}	response.ApiResponse
//	@Response		500		{object}	response.ApiResponse
//	@Router			/v1/idstar/training/update [put]
func (ctrl *TrainingController) UpdateTraining(ctx *gin.Context) {
	req := request.TrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.UpdateTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetTrainingById godoc
//
//	@Summary		Get Training By Id
//	@Description	To authorize this API, get token from user-login
//	@Id				GetTrainingById
//	@Tags			idstar/training
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Param			id	path		int	true	"Training ID"
//	@Response		200	{object}	response.ApiResponse
//	@Response		400	{object}	response.ApiResponse
//	@Response		404	{object}	response.ApiResponse
//	@Response		500	{object}	response.ApiResponse
//	@Router			/v1/idstar/training/{id} [get]
func (ctrl *TrainingController) GetTrainingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetTrainingById(id)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// GetTrainingList godoc
//
//	@Summary		Get Training List
//	@Description	To authorize this API, get token from user-login
//	@Id				GetTrainingList
//	@Tags			idstar/training
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Param			page		query		int		false	"Page"
//	@Param			size		query		int		false	"Size"
//	@Param			field		query		string	false	"Field"
//	@Param			direction	query		string	false	"Direction"
//	@Response		200			{object}	response.ApiResponse
//	@Response		400			{object}	response.ApiResponse
//	@Response		500			{object}	response.ApiResponse
//	@Router			/v1/idstar/training/list [get]
func (ctrl *TrainingController) GetTrainingList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetTrainingList(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// DeleteTraining godoc
//
//	@Summary		Delete Training
//	@Description	To authorize this API, get token from user-login
//	@Id				DeleteTraining
//	@Tags			idstar/training
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Param			request	body		request.IdRequest	true	"Delete Training Request"
//	@Response		200		{object}	response.ApiResponse
//	@Response		400		{object}	response.ApiResponse
//	@Response		500		{object}	response.ApiResponse
//	@Router			/v1/idstar/training/delete [delete]
func (ctrl *TrainingController) DeleteTraining(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	code, err := ctrl.svc.DeleteTraining(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponseNoData(ctx)
}
