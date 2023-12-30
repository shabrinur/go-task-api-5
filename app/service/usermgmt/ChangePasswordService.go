package service

import (
	"fmt"
	"idstar-idp/rest-api/app/dto"
	"idstar-idp/rest-api/app/dto/request/login"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"
)

type ChangePasswordService struct {
	OtpService
	mailUtil util.MailUtil
	pwdUtil  util.PasswordUtil
}

func NewChangePasswordService(repo repository.UserRepository, userMgmtUtil util.UserMgmtUtil) *ChangePasswordService {
	changePasswordService := &ChangePasswordService{
		mailUtil: userMgmtUtil.MailUtil,
		pwdUtil:  util.PasswordUtil{},
	}
	changePasswordService.OtpService = *NewOtpService(repo, userMgmtUtil.OtpUtil)
	return changePasswordService
}

func (svc *ChangePasswordService) GetChangePasswordOtp(req login.OtpRequest) (interface{}, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	user, code, err := svc.getExistingUserData(req.Username, true)
	if err != nil {
		return nil, code, err
	}

	otp, otpExpiredOn, err := svc.saveUserOtp(user, false)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	err = svc.mailUtil.SendChangePasswordMail(user.Username, user.Fullname, otp, otpExpiredOn)
	if err != nil {
		otpInfo := &dto.OtpInfo{
			Otp:       otp,
			ExpiredOn: otpExpiredOn,
		}
		backupResponse := &dto.OtpBackupData{
			Info:    fmt.Sprint("Request OTP success for user ", req.Username, "; please use the provided OTP for password reset"),
			Error:   err.Error(),
			OtpInfo: *otpInfo,
		}
		return backupResponse, 0, nil
	}
	return fmt.Sprint("Request OTP success for user ", req.Username, "; please check your mail for password reset instruction"), 0, nil
}

func (svc *ChangePasswordService) ValidatePasswordOtp(req login.OtpRequest) (int, error) {
	err := req.Validate(true)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user, code, err := svc.getExistingUserData(req.Username, true)
	if err != nil {
		return code, err
	}

	err = svc.otpUtil.ValidateOtp(req.Otp, user.Otp, user.OtpExpiredDate)
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

	user, code, err := svc.getExistingUserData(req.Username, true)
	if err != nil {
		return code, err
	}

	err = svc.otpUtil.ValidateOtp(req.Otp, user.Otp, user.OtpExpiredDate)
	if err != nil {
		return http.StatusBadRequest, err
	}

	newPwd, err := svc.pwdUtil.Encrypt(req.NewPassword)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.Password = *newPwd
	_, err = svc.repo.UpdatePassword(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, nil
}
