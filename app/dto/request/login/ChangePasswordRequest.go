package login

import (
	"github.com/go-playground/validator/v10"
)

type ChangePasswordRequest struct {
	Username        string `json:"username" validate:"required,email"`
	Otp             string `json:"otp" validate:"required,alphanum"`
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=16"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=NewPassword"`
}

func (c *ChangePasswordRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
