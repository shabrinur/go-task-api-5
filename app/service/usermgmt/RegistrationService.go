package service

import (
	"database/sql"
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto"
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	model "idstar-idp/rest-api/app/model/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"
)

type RegistrationService struct {
	OtpService
	roleModuleRepo repository.RoleModuleRepository
	mailUtil       util.MailUtil
	pwdUtil        util.PasswordUtil
}

func NewRegistrationService(userMgmtRepo repository.UserMgmtRepository, userMgmtUtil util.UserMgmtUtil) *RegistrationService {
	registrationService := &RegistrationService{
		roleModuleRepo: userMgmtRepo.RoleModuleRepository,
		mailUtil:       userMgmtUtil.MailUtil,
		pwdUtil:        util.PasswordUtil{},
	}
	registrationService.OtpService = *NewOtpService(userMgmtRepo.UserRepository, userMgmtUtil.OtpUtil)
	return registrationService
}

func (svc *RegistrationService) CreateUser(req login.LoginUserPassRequest) (*rsdata.RegistrationData, int, error) {
	err := req.Validate(true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if svc.repo.CheckUserAlreadyExists(req.Username) {
		return nil, http.StatusConflict, errors.New(fmt.Sprint("user ", req.Username, " already exists"))
	}

	role, err := svc.roleModuleRepo.GetDefaultUserRole()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	permissions, err := svc.roleModuleRepo.GetPermissions(role.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	password, err := svc.pwdUtil.Encrypt(req.Password)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	dbObj := &model.UserModel{
		Fullname: req.Name,
		Username: req.Username,
		Password: *password,
		IDRole:   role.ID,
	}
	otp, otpExpiredOn, err := svc.saveUserOtp(dbObj, true)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	info := &rsdata.OtpBackupData{}
	err = svc.mailUtil.SendUserActivationMail(req.Username, req.Name, otp, otpExpiredOn)
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

	userPermission := &dto.UserPermission{
		Role:        role.Name,
		Permissions: permissions,
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
	return regData, 0, nil
}

func (svc *RegistrationService) GetActivationLink(req login.OtpRequest) (interface{}, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	user, code, err := svc.getExistingUserData(req.Username, false)
	if err != nil {
		return nil, code, err
	}

	otp, otpExpiredOn, err := svc.saveUserOtp(user, false)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	err = svc.mailUtil.SendUserActivationMail(user.Username, user.Fullname, otp, otpExpiredOn)
	if err != nil {
		otpInfo := &dto.OtpInfo{
			Otp:       otp,
			ExpiredOn: otpExpiredOn,
		}
		backupResponse := &rsdata.OtpBackupData{
			Info:    fmt.Sprint("Activation code request success for user ", user.Username, "; please use the provided code to proceed"),
			Error:   err.Error(),
			OtpInfo: *otpInfo,
		}
		return backupResponse, 0, nil
	}
	return fmt.Sprint("Activation link request success for user ", user.Username, "; please check your mail for further instruction"), 0, nil
}

func (svc *RegistrationService) activateUser(username string, otp string) (string, int, error) {
	user, code, err := svc.getExistingUserData(username, false)
	if err != nil {
		return "", code, err
	}

	if user.AccountActivated.Bool {
		return fmt.Sprint("User ", username, " already active! Please proceed to login"), 0, nil
	}

	err = svc.otpUtil.ValidateOtp(otp, user.Otp, user.OtpExpiredDate)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	user.AccountActivated = sql.NullBool{Valid: true, Bool: true}
	_, err = svc.repo.ActivateUser(user)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return fmt.Sprint("Congratulation, ", username, " is now active! Please proceed to login"), 0, nil
}

func (svc *RegistrationService) ActivateByCode(req login.OtpRequest) (string, int, error) {
	err := req.Validate(true)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	return svc.activateUser(req.Username, req.Otp)
}

func (svc *RegistrationService) ActivateByLink(encodedString string) (string, int, error) {
	username, otp, err := util.DecodeFromActivationLink(encodedString)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	return svc.activateUser(username, otp)
}
