package service

import (
	"database/sql"
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto"
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"idstar-idp/rest-api/app/service/usermgmt/helper"
	"idstar-idp/rest-api/app/util"
	"net/http"
)

type RegistrationService struct {
	helper.UserHelper
}

func NewRegistrationService(userHelper helper.UserHelper) *RegistrationService {
	return &RegistrationService{
		UserHelper: userHelper,
	}
}

func (svc *RegistrationService) RegisterUser(req login.LoginUserPassRequest) (*rsdata.RegistrationData, int, error) {
	err := req.Validate(true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if svc.GetUserRepo().CheckUserAlreadyExists(req.Username) {
		return nil, http.StatusConflict, errors.New(fmt.Sprint("user ", req.Username, " already exists"))
	}

	dbObj, userPermission, err := svc.SetUserRoleAndPermissions(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	password, err := svc.GetPwdUtil().Encrypt(req.Password)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	dbObj.Password = *password

	otp, otpExpiredOn, err := svc.SaveUserOtp(dbObj, true)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return svc.SendInitialActivationMail(req, userPermission, otp, otpExpiredOn), 0, nil
}

func (svc *RegistrationService) GetActivationLink(req login.OtpRequest) (interface{}, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	user, code, err := svc.GetExistingUserData(req.Username, false)
	if err != nil {
		return nil, code, err
	}

	otp, otpExpiredOn, err := svc.SaveUserOtp(user, false)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	err = svc.GetMailUtil().SendUserActivationMail(user.Username, user.Fullname, otp, otpExpiredOn)
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
	user, code, err := svc.GetExistingUserData(username, false)
	if err != nil {
		return "", code, err
	}

	if user.AccountActivated.Bool {
		return fmt.Sprint("User ", username, " already active! Please proceed to login"), 0, nil
	}

	err = svc.GetOtpUtil().ValidateOtp(otp, user.Otp, user.OtpExpiredDate)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	user.AccountActivated = sql.NullBool{Valid: true, Bool: true}
	_, err = svc.GetUserRepo().ActivateUser(user)
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
