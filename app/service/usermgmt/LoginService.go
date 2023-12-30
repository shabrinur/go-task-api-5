package service

import (
	"errors"
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"idstar-idp/rest-api/app/service/usermgmt/helper"
	"net/http"
	"strings"
)

type LoginService struct {
	helper.UserHelper
}

func NewLoginService(userHelper helper.UserHelper) *LoginService {
	return &LoginService{
		UserHelper: userHelper,
	}
}

func (svc *LoginService) LoginUserPassword(req login.LoginUserPassRequest) (*rsdata.LoginData, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, code, err := svc.FindExistingUser(req, false)
	if err != nil {
		return nil, code, err
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
