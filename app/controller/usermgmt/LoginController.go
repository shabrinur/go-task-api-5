package controller

import (
	"encoding/json"
	"idstar-idp/rest-api/app/dto/request/login"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	svc *service.LoginService
}

func NewLoginController(svc *service.LoginService) *LoginController {
	return &LoginController{svc}
}

// Login godoc
//
//	@Summary	Login
//	@Id			Login
//	@Tags		usermanagement
//	@Accept		json
//	@Produce	json
//	@Param		request	body		login.LoginUserPassRequest	true	"Login Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/user-login/login [post]
func (ctrl *LoginController) Login(ctx *gin.Context) {
	bodyAsByteArray, _ := io.ReadAll(ctx.Request.Body)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyAsByteArray, &jsonMap)
	if err != nil {
		util.SetErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	if _, ok := jsonMap["password"]; ok {
		req := login.LoginUserPassRequest{}
		err := json.Unmarshal(bodyAsByteArray, &req)
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
}
