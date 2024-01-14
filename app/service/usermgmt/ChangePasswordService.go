package service

import (
	"fmt"
	"idstar-idp/rest-api/app/dto"
	"idstar-idp/rest-api/app/dto/request/login"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"idstar-idp/rest-api/app/service/usermgmt/helper"
	"net/http"
)

type ChangePasswordService struct {
	helper.OtpHelper
}

func NewChangePasswordService(otpHelper helper.OtpHelper) *ChangePasswordService {
	return &ChangePasswordService{
		OtpHelper: otpHelper,
	}
}

func (svc *ChangePasswordService) GetChangePasswordOtp(req login.OtpRequest) (interface{}, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	user, code, err := svc.GetExistingUserData(req.Username, true)
	if err != nil {
		return nil, code, err
	}

	otp, otpExpiredOn, err := svc.SaveUserOtp(user, false)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	err = svc.GetMailUtil().SendChangePasswordMail(user.Username, user.Fullname, otp, otpExpiredOn)
	if err != nil {
		otpInfo := &dto.OtpInfo{
			Otp:       otp,
			ExpiredOn: &otpExpiredOn,
		}
		backupResponse := &rsdata.OtpBackupData{
			Info:    fmt.Sprint("Password reset request success for user ", req.Username, "; please use the provided code to proceed"),
			Error:   err.Error(),
			OtpInfo: *otpInfo,
		}
		return backupResponse, 0, nil
	}
	return fmt.Sprint("Password reset request success for user ", req.Username, "; please check your mail for further instruction"), 0, nil
}

func (svc *ChangePasswordService) ValidatePasswordOtp(req login.OtpRequest) (int, error) {
	err := req.Validate(true)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user, code, err := svc.GetExistingUserData(req.Username, true)
	if err != nil {
		return code, err
	}

	err = svc.GetOtpUtil().ValidateOtp(req.Otp, user.Otp, user.OtpExpiredDate)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return 0, nil
}

func (svc *ChangePasswordService) ChangePassword(req login.ChangePasswordRequest) (int, error) {
	err := req.Validate()
	if err != nil {
		return http.StatusBadRequest, err
	}

	user, code, err := svc.GetExistingUserData(req.Username, true)
	if err != nil {
		return code, err
	}

	err = svc.GetOtpUtil().ValidateOtp(req.Otp, user.Otp, user.OtpExpiredDate)
	if err != nil {
		return http.StatusBadRequest, err
	}

	newPwd, err := svc.GetPwdUtil().Encrypt(req.NewPassword)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.Password = *newPwd
	_, err = svc.GetUserRepo().UpdatePassword(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, nil
}
