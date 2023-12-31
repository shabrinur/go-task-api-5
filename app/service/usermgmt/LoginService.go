package service

import (
	"database/sql"
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"idstar-idp/rest-api/app/service/usermgmt/helper"
	"idstar-idp/rest-api/app/service/usermgmt/oauth"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/exp/slices"
)

type LoginService struct {
	helper.UserHelper
	google *oauth.OauthGoogle
}

func NewLoginService(userHelper helper.UserHelper) *LoginService {
	return &LoginService{
		UserHelper: userHelper,
		google:     oauth.InitOauthGoogle(),
	}
}

func (svc *LoginService) LoginUserPassword(req login.RegistrationLoginRequest) (*rsdata.LoginData, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, code, err := svc.FindExistingUser(req, false)
	if err != nil {
		return nil, code, err
	}

	if result.Oauth.Bool && result.Password == "" {
		return nil, http.StatusUnauthorized, errors.New(fmt.Sprint("password has not been set; please login with " + result.OauthProvider + " Oauth"))
	}

	inputPwd := strings.TrimSpace(req.Password)
	savedPwd, err := svc.GetPwdUtil().Decrypt(result.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if !strings.EqualFold(inputPwd, *savedPwd) {
		return nil, http.StatusBadRequest, errors.New("invalid password")
	}

	return svc.GetLoginData(result, "password")
}

func (svc *LoginService) ValidateOauthProvider(provider string) error {
	if provider == "" {
		return errors.New("oauth provider missing")
	}
	validProviders := []string{"google"}
	if !slices.Contains(validProviders, strings.ToLower(provider)) {
		return errors.New(fmt.Sprint("invalid oauth provider; acceptable provider: ", strings.Join(validProviders, ", ")))
	}
	return nil
}

func (svc *LoginService) GetProviderAuthAddress(provider string) string {
	switch provider {
	case svc.google.OauthPath:
		return svc.google.GetGoogleAuthAddress()
	}
	return ""
}

func (svc *LoginService) OauthLogin(provider string, queryMap url.Values) (interface{}, int, error) {
	switch provider {
	case svc.google.OauthPath:
		return svc.googleOauthLogin(queryMap)
	}
	return "", 0, nil

}

func (svc *LoginService) googleOauthLogin(queryMap url.Values) (interface{}, int, error) {
	authCode := queryMap.Get("code")
	if authCode == "" {
		return nil, http.StatusBadRequest, errors.New("authorization code missing")
	}
	token, err := svc.google.GetGoogleOauthToken(authCode)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	loginReq, err := svc.google.GetGoogleUserInfo(token.AccessToken, token.IdToken)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	result, code, err := svc.FindExistingUser(*loginReq, true)
	if err != nil {
		return nil, code, err
	}

	if result != nil {
		if result.OauthProvider != "" && !strings.EqualFold(result.OauthProvider, svc.google.ProviderName) {
			return nil, http.StatusUnauthorized, errors.New(fmt.Sprint("invalid oauth login for user ", loginReq.Username, "; please login with ", result.OauthProvider, " Oauth"))
		}
		result.AccessToken = token.AccessToken
		result.AccessTokenExpiredDate = token.ExpireIn
		if result.Oauth.Bool {
			_, err = svc.GetUserRepo().UpdateOauthInfo(result)
		} else {
			result.Fullname = loginReq.Name
			result.Oauth = sql.NullBool{Valid: true, Bool: true}
			result.OauthProvider = svc.google.ProviderName
			_, err = svc.GetUserRepo().EnableOauthLogin(result)
		}
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		return svc.GetLoginData(result, svc.google.OauthPath)
	}

	result, userPermission, err := svc.SetUserRoleAndPermissions(*loginReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	result.AccessToken = token.AccessToken
	result.AccessTokenExpiredDate = token.ExpireIn
	result.Oauth = sql.NullBool{Valid: true, Bool: true}
	result.OauthProvider = svc.google.ProviderName

	otp, otpExpiredOn, err := svc.SaveUserOtp(result, true)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return svc.SendInitialActivationMail(*loginReq, userPermission, otp, otpExpiredOn), 0, nil
}
