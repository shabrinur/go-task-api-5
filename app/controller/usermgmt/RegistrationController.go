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

// CreateUser godoc
//
//	@Summary	Create User
//	@Id			CreateUser
//	@Tags		user-register
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.LoginUserPassRequest	true	"Create User Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/user-register [post]
func (ctrl *RegistrationController) CreateUser(ctx *gin.Context) {
	req := login.LoginUserPassRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.CreateUser(req)
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
//	@Tags		user-register
//	@Accept		json
//	@Produce	json
//	@Param		go	query		string	true	"Encoded Activation Parameter"
//	@Response	200	{object}	response.ApiResponse
//	@Response	400	{object}	response.ApiResponse
//	@Response	500	{object}	response.ApiResponse
//	@Router		/v1/user-register/send-link [post]
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

// ActivateByCode godoc
//
//	@Summary	Activate By Code
//	@Id			ActivateByCode
//	@Tags		user-register
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.OtpRequest	true	"Activation Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/user-register/activate [post]
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

// ActivateByLink godoc
//
//	@Summary	Activate By Link
//	@Id			ActivateByLink
//	@Tags		user-register
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.OtpRequest	true	"Activation Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/user-register/activate [get]
func (ctrl *RegistrationController) ActivateByLink(ctx *gin.Context) {
	encodedString := ctx.Query("go")
	if encodedString == "" {
		util.SetErrorResponse(ctx, errors.New("activation code missing"), http.StatusBadRequest)
		return
	}

	msg, code, err := ctrl.svc.ActivateByLink(encodedString)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, msg)
}
