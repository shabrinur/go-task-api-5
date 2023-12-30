package service

import (
	"errors"
	"fmt"
	model "idstar-idp/rest-api/app/model/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"time"
)

type OtpService struct {
	repo    repository.UserRepository
	otpUtil util.OtpUtil
}

func NewOtpService(repo repository.UserRepository, otpUtil util.OtpUtil) *OtpService {
	return &OtpService{
		repo:    repo,
		otpUtil: otpUtil,
	}
}

func (svc *OtpService) getExistingUserData(username string, activatedRequired bool) (*model.UserModel, int, error) {
	result, err := svc.repo.GetByUsername(username)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("user ", username, " not registered"))
	}
	if activatedRequired {
		if !result.AccountActivated.Valid || !result.AccountActivated.Bool {
			return nil, http.StatusUnauthorized, errors.New(fmt.Sprint("user ", username, " not yet activated; please check your mail for activation link"))
		}
	}
	return result, 0, nil
}

func (svc *OtpService) saveUserOtp(user *model.UserModel, isNewUser bool) (string, time.Time, error) {
	otp, otpExpiredOn := svc.otpUtil.GenerateOtp()

	user.Otp = otp
	user.OtpExpiredDate = otpExpiredOn

	if isNewUser {
		_, err := svc.repo.Create(user)
		if err != nil {
			return "", time.Time{}, err
		}
	} else {
		_, err := svc.repo.UpdateOtp(user)
		if err != nil {
			return "", time.Time{}, err
		}
	}

	return otp, otpExpiredOn, nil
}
