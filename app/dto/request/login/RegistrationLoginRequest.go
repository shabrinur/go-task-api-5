package login

import (
	"github.com/go-playground/validator/v10"
)

type RegistrationLoginRequest struct {
	Name            string `json:"name,omitempty" validate:"required"`
	Username        string `json:"username" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=16"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

func (c *RegistrationLoginRequest) Validate(isRegistration bool) error {
	if !isRegistration {
		c.Name = "dummy"
		c.ConfirmPassword = c.Password
	}

	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
