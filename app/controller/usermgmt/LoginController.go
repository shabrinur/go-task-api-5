package controller

import (
	"idstar-idp/rest-api/app/dto/request/login"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	svc *service.LoginService
}

func NewLoginController(svc *service.LoginService) *LoginController {
	return &LoginController{svc}
}

// UserPassLogin godoc
//
//	@Summary	Username & Password Login
//	@Id			UserPassLogin
//	@Tags		user-login
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.LoginUserPassRequest	true	"Username & Password Login Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/user-login/login [post]
func (ctrl *LoginController) UserPassLogin(ctx *gin.Context) {
	req := login.LoginUserPassRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	result, code, err := ctrl.svc.LoginUserPassword(req)
	if err != nil {
		util.SetErrorResponse(ctx, err, code)
		return
	}
	util.SetSuccessResponse(ctx, result)
}
