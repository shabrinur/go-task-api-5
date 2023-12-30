package login

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type OtpRequest struct {
	Username string `json:"username" validate:"required,email"`
	Otp      string `json:"otp,omitempty" validate:"alphanum"`
}

func (c *OtpRequest) Validate(isValidateOtp bool) error {
	if isValidateOtp {
		if c.Otp == "" {
			return errors.New("value for field 'Otp' is required")
		}
	} else {
		c.Otp = "dummy"
	}

	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
