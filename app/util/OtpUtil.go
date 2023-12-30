package util

import (
	"errors"
	"idstar-idp/rest-api/app/config"
	"math/rand"
	"strings"
	"time"
)

type OtpUtil struct {
	random     *rand.Rand
	otpLength  int
	otpCharset string
	otpExpire  int
}

func NewOtpUtil() *OtpUtil {
	return &OtpUtil{
		random:     rand.New(rand.NewSource(time.Now().UnixNano())),
		otpLength:  config.GetConfigIntValue("otp.length"),
		otpCharset: config.GetConfigValue("otp.charset"),
		otpExpire:  config.GetConfigIntValue("otp.expire.ms")}
}

func (c *OtpUtil) GenerateOtp() (string, time.Time) {
	result := make([]byte, c.otpLength)
	for i := range result {
		result[i] = c.otpCharset[c.random.Intn(len(c.otpCharset))]
	}

	otpExpiredDate := time.Now().Add(time.Millisecond * time.Duration(int64(c.otpExpire)))
	return string(result), otpExpiredDate
}

func (c *OtpUtil) ValidateOtp(inputOtp string, savedOtp string, otpExpiredDate time.Time) error {
	if !strings.EqualFold(inputOtp, savedOtp) {
		return errors.New("invalid OTP; please re-check input")
	}
	if !time.Now().Before(otpExpiredDate) {
		return errors.New("expired OTP; please request new OTP")
	}
	return nil
}
