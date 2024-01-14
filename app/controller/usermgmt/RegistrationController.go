package controller

import (
	"errors"
	"idstar-idp/rest-api/app/dto/request/login"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	svc *service.RegistrationService
}

func NewRegistrationController(svc *service.RegistrationService) *RegistrationController {
	return &RegistrationController{svc}
}

// RegisterUser godoc
//
//	@Summary	Register User
//	@Id			RegisterUser
//	@Tags		registration
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.RegistrationLoginRequest	true	"Register User Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/registration [post]
func (ctrl *RegistrationController) RegisterUser(ctx *gin.Context) {
	req := login.RegistrationLoginRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.RegisterUser(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}
	util.SetSuccessResponse(ctx, result)
}

// GetActivationLink godoc
//
//	@Summary	Get Activation Link
//	@Id			GetActivationLink
//	@Tags		registration
//	@Accept		json
//	@Produce	json
//	@Param		go	query		string	true	"Encoded Activation Parameter"
//	@Response	200	{object}	response.ApiResponse
//	@Response	400	{object}	response.ApiResponse
//	@Response	500	{object}	response.ApiResponse
//	@Router		/v1/registration/send-link [post]
func (ctrl *RegistrationController) GetActivationLink(ctx *gin.Context) {
	req := login.OtpRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetActivationLink(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// ActivateByLink godoc
//
//	@Summary		Activate By Link
//	@Description	This API produces text/html. It is preferable to execute the API call from web browser for testing.
//	@Id				ActivateByLink
//	@Tags			registration
//	@Accept			json
//	@Produce		text/html
//	@Param			go	query		string	true	"Encoded Activation Request"
//	@Response		200	{string}	string	"text/html"
//	@Response		400	{string}	string	"text/html"
//	@Response		500	{string}	string	"text/html"
//	@Router			/v1/registration/activate [get]
func (ctrl *RegistrationController) ActivateByLink(ctx *gin.Context) {
	encodedString := ctx.Query("go")
	if encodedString == "" {
		util.ShowErrorResponsePage(ctx, errors.New("activation code missing"), http.StatusBadRequest)
		return
	}

	msg, code, err := ctrl.svc.ActivateByLink(encodedString)
	if err != nil {
		util.ShowErrorResponsePage(ctx, err, code)
		return
	}

	util.ShowActivationResponsePage(ctx, msg)
}

// ActivateByCode godoc
//
//	@Summary	Activate By Code
//	@Id			ActivateByCode
//	@Tags		registration
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.OtpRequest	true	"Activation Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/registration/activate [post]
func (ctrl *RegistrationController) ActivateByCode(ctx *gin.Context) {
	req := login.OtpRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	msg, code, err := ctrl.svc.ActivateByCode(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, msg)
}
