package helper

import (
	"errors"
	"fmt"
	model "idstar-idp/rest-api/app/model/usermgmt"
	repository "idstar-idp/rest-api/app/repository/usermgmt"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"time"
)

type OtpHelper struct {
	userRepo     repository.UserRepository
	userMgmtUtil util.UserMgmtUtil
}

func NewOtpHelper(userRepo repository.UserRepository, userMgmtUtil util.UserMgmtUtil) *OtpHelper {
	return &OtpHelper{
		userRepo:     userRepo,
		userMgmtUtil: userMgmtUtil,
	}
}

func (h *OtpHelper) GetUserRepo() *repository.UserRepository {
	return &h.userRepo
}

func (h *OtpHelper) GetOtpUtil() *util.OtpUtil {
	return &h.userMgmtUtil.OtpUtil
}

func (h *OtpHelper) GetMailUtil() *util.MailUtil {
	return &h.userMgmtUtil.MailUtil
}

func (h *OtpHelper) GetPwdUtil() *util.PasswordUtil {
	return &h.userMgmtUtil.PwdUtil
}

func (h *OtpHelper) GetExistingUserData(username string, activatedRequired bool) (*model.UserModel, int, error) {
	result, err := h.userRepo.GetByUsername(username)
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

func (h *OtpHelper) SaveUserOtp(user *model.UserModel, isNewUser bool) (string, time.Time, error) {
	otp, otpExpiredOn := h.userMgmtUtil.OtpUtil.GenerateOtp()

	user.Otp = otp
	user.OtpExpiredDate = otpExpiredOn

	if isNewUser {
		_, err := h.userRepo.Create(user)
		if err != nil {
			return "", time.Time{}, err
		}
	} else {
		_, err := h.userRepo.UpdateOtp(user)
		if err != nil {
			return "", time.Time{}, err
		}
	}

	return otp, otpExpiredOn, nil
}
