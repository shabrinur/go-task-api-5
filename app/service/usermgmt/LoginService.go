package service

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto"
	"idstar-idp/rest-api/app/dto/request/login"
	model "idstar-idp/rest-api/app/model/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"strings"
	"time"
)

type LoginService struct {
	userMgmtRepo       repository.UserMgmtRepository
	pwdUtil            util.PasswordUtil
	authTokenIssuer    string
	authTokenAudiences string
	authTokenKey       string
	authTokenExpire    int
}

func NewLoginService(userMgmtRepo repository.UserMgmtRepository) *LoginService {
	return &LoginService{
		userMgmtRepo:       userMgmtRepo,
		pwdUtil:            util.PasswordUtil{},
		authTokenIssuer:    config.GetConfigValue("authtoken.issuer"),
		authTokenAudiences: config.GetConfigValue("authtoken.audience"),
		authTokenKey:       config.GetConfigValue("authtoken.secret"),
		authTokenExpire:    config.GetConfigIntValue("authtoken.expire.ms")}
}

func (svc *LoginService) getLoginData(user *model.UserModel, authMethod string) (*dto.LoginData, int, error) {
	permissions, err := svc.userMgmtRepo.GetPermissions(user.Role.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tokenPayload := &dto.AuthTokenPayload{
		Username:    user.Username,
		AuthMethod:  authMethod,
		RoleType:    user.Role.RoleType,
		Permissions: permissions,
	}

	duration := time.Millisecond * time.Duration(int64(svc.authTokenExpire))
	tokenInfo, err := util.GenerateToken(duration, *tokenPayload, svc.authTokenIssuer, svc.authTokenAudiences, svc.authTokenKey)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	role := &dto.UserPermission{
		Role:        user.Role.Name,
		Permissions: permissions,
	}
	loginData := &dto.LoginData{
		Username:       user.Username,
		Fullname:       user.Fullname,
		TokenInfo:      *tokenInfo,
		UserPermission: *role,
	}

	return loginData, 0, nil
}

func (svc *LoginService) LoginUserPassword(req login.LoginUserPassRequest) (*dto.LoginData, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := svc.userMgmtRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("user ", req.Username, " not registered"))
	}

	if !result.AccountActivated.Valid || !result.AccountActivated.Bool {
		return nil, http.StatusUnauthorized, errors.New(fmt.Sprint("user ", req.Username, " not yet activated; please check your mail for activation link"))
	}

	inputPwd := strings.TrimSpace(req.Password)
	savedPwd, err := svc.pwdUtil.Decrypt(result.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if !strings.EqualFold(inputPwd, *savedPwd) {
		return nil, http.StatusBadRequest, errors.New("invalid password")
	}

	return svc.getLoginData(result, "password")
}
