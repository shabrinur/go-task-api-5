package controller

import (
	"idstar-idp/rest-api/app/dto/request/login"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangePasswordController struct {
	svc *service.ChangePasswordService
}

func NewChangePasswordController(svc *service.ChangePasswordService) *ChangePasswordController {
	return &ChangePasswordController{svc}
}

// GetChangePasswordOtp godoc
//
//	@Summary	Get Change Password Otp
//	@Id			GetChangePasswordOtp
//	@Tags		usermanagement
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.OtpRequest	true	"Change Password OTP Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/forget-password/send-otp [post]
func (ctrl *ChangePasswordController) GetChangePasswordOtp(ctx *gin.Context) {
	req := login.OtpRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.GetChangePasswordOtp(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponse(ctx, result)
}

// ValidateChangePasswordOtp godoc
//
//	@Summary	Validate Change Password Otp
//	@Id			ValidateChangePasswordOtp
//	@Tags		usermanagement
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.OtpRequest	true	"Change Password OTP Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/forget-password/validate-otp [post]
func (ctrl *ChangePasswordController) ValidateChangePasswordOtp(ctx *gin.Context) {
	req := login.OtpRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	code, err := ctrl.svc.ValidatePasswordOtp(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponseNoData(ctx)
}

// ChangePassword godoc
//
//	@Summary	Change Password
//	@Id			ChangePassword
//	@Tags		usermanagement
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.ChangePasswordRequest	true	"Change Password Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/forget-password/change-password [post]
func (ctrl *ChangePasswordController) ChangePassword(ctx *gin.Context) {
	req := login.ChangePasswordRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	code, err := ctrl.svc.ChangePassword(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}

	util.SetSuccessResponseNoData(ctx)
}
