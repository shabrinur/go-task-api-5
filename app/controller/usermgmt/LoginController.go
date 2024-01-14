package controller

import (
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	service "idstar-idp/rest-api/app/service/usermgmt"
	"idstar-idp/rest-api/app/util"
	"log"
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
//	@Param		request	body		login.RegistrationLoginRequest	true	"Username & Password Login Request"
//	@Response	200		{object}	response.ApiResponse
//	@Response	400		{object}	response.ApiResponse
//	@Response	500		{object}	response.ApiResponse
//	@Router		/v1/user-login/login [post]
func (ctrl *LoginController) UserPassLogin(ctx *gin.Context) {
	req := login.RegistrationLoginRequest{}

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

// OauthLogin godoc
//
//	@Summary		Oauth Login
//	@Description	This API contains redirect and produces text/html. It is preferable to execute the API call from web browser for testing.
//	@Id				OauthLogin
//	@Tags			user-login
//	@Accept			json
//	@Produce		text/html
//	@Param			provider	path		string	true	"Oauth Provider"
//	@Response		200			{string}	string	"text/html"
//	@Response		302			{string}	string	"text/html"
//	@Response		400			{string}	string	"text/html"
//	@Response		500			{string}	string	"text/html"
//	@Router			/v1/user-login/oauth/{provider} [get]
func (ctrl *LoginController) OauthLogin(ctx *gin.Context) {
	oauthProvider := ctx.Param("provider")
	err := ctrl.svc.ValidateOauthProvider(oauthProvider)
	if err != nil {
		util.ShowErrorResponsePage(ctx, err, http.StatusBadRequest)
		return
	}

	queryMap := ctx.Request.URL.Query()
	if len(queryMap) == 0 {
		authAddress := ctrl.svc.GetProviderAuthAddress(oauthProvider)
		log.Printf("accessing %s oauth address at %s", oauthProvider, authAddress)
		ctx.Redirect(http.StatusFound, authAddress)
		return
	}

	result, code, err := ctrl.svc.OauthLogin(oauthProvider, queryMap)
	if err != nil {
		util.ShowErrorResponsePage(ctx, err, code)
		return
	}
	data, ok := result.(*rsdata.LoginData)
	if !ok {
		util.ShowRegistrationResponsePage(ctx, result.(*rsdata.RegistrationData))
		return
	}
	util.ShowLoginResponsePage(ctx, data)
}
