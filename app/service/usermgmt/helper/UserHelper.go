package helper

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto"
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	model "idstar-idp/rest-api/app/model/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"time"
)

type UserHelper struct {
	OtpHelper
	repo               repository.RoleModuleRepository
	authTokenIssuer    string
	authTokenAudiences string
	authTokenKey       string
	authTokenExpire    int
}

func NewUserHelper(repo repository.RoleModuleRepository, otpHelper OtpHelper) *UserHelper {
	return &UserHelper{
		OtpHelper:          otpHelper,
		repo:               repo,
		authTokenIssuer:    config.GetConfigValue("authtoken.issuer"),
		authTokenAudiences: config.GetConfigValue("authtoken.audience"),
		authTokenKey:       config.GetConfigValue("authtoken.secret"),
		authTokenExpire:    config.GetConfigIntValue("authtoken.expire.ms")}
}

func (h *UserHelper) SetUserRoleAndPermissions(req login.RegistrationLoginRequest) (*model.UserModel, *dto.UserPermission, error) {
	role, err := h.repo.GetDefaultUserRole()
	if err != nil {
		return nil, nil, err
	}

	permissions, err := h.repo.GetPermissions(role.ID)
	if err != nil {
		return nil, nil, err
	}

	user := &model.UserModel{
		Fullname: req.Name,
		Username: req.Username,
		IDRole:   role.ID,
	}
	userPermission := &dto.UserPermission{
		Role:        role.Name,
		Permissions: permissions,
	}
	return user, userPermission, nil
}

func (h *UserHelper) SendInitialActivationMail(req login.RegistrationLoginRequest, userPermission *dto.UserPermission, otp string, otpExpiredOn time.Time) *rsdata.RegistrationData {
	info := &rsdata.OtpBackupData{}
	err := h.GetMailUtil().SendUserActivationMail(req.Username, req.Name, otp, otpExpiredOn)
	if err != nil {
		otpInfo := &dto.OtpInfo{
			Otp:       otp,
			ExpiredOn: otpExpiredOn,
		}
		info.Info = fmt.Sprint("Reqistration success for user ", req.Username, "; please use the provided code for activation")
		info.OtpInfo = *otpInfo
		info.Error = err.Error()
	} else {
		info.Info = fmt.Sprint("Reqistration success for user ", req.Username, "; please check your mail for further instruction")
	}

	userInfo := &dto.UserInfo{
		Fullname:       req.Name,
		Username:       req.Username,
		IsUserActive:   false,
		UserPermission: *userPermission,
	}
	regData := &rsdata.RegistrationData{
		UserInfo:      *userInfo,
		OtpBackupData: *info,
	}
	return regData
}

func (h *UserHelper) FindExistingUser(req login.RegistrationLoginRequest, isOauth bool) (*model.UserModel, int, error) {
	result, err := h.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		if isOauth {
			return nil, 0, nil
		} else {
			return nil, http.StatusNotFound, errors.New(fmt.Sprint("user ", req.Username, " not registered"))
		}
	}

	if !result.AccountActivated.Valid || !result.AccountActivated.Bool {
		return nil, http.StatusUnauthorized, errors.New(fmt.Sprint("user ", req.Username, " not yet activated; please check your mail for activation link"))
	}
	return result, 0, nil
}

func (h *UserHelper) GetLoginData(user *model.UserModel, authMethod string) (*rsdata.LoginData, int, error) {
	permissions, err := h.repo.GetPermissions(user.Role.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tokenPayload := &dto.AuthTokenPayload{
		Username:    user.Username,
		AuthMethod:  authMethod,
		RoleType:    user.Role.RoleType,
		Permissions: permissions,
	}

	duration := time.Millisecond * time.Duration(int64(h.authTokenExpire))
	tokenInfo, err := util.GenerateToken(duration, *tokenPayload, h.authTokenIssuer, h.authTokenAudiences, h.authTokenKey)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	role := &dto.UserPermission{
		Role:        user.Role.Name,
		Permissions: permissions,
	}
	userInfo := &dto.UserInfo{
		Fullname:       user.Fullname,
		Username:       user.Username,
		IsUserActive:   user.AccountActivated.Bool,
		UserPermission: *role,
	}
	loginData := &rsdata.LoginData{
		UserInfo:  *userInfo,
		TokenInfo: *tokenInfo,
	}

	return loginData, 0, nil
}
