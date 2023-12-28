package login

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type LoginUserPassRequest struct {
	Name     string `json:"name,omitempty" validate:"alpha"`
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

func (c *LoginUserPassRequest) Validate(isRegistration bool) error {
	if isRegistration {
		if c.Name == "" {
			return errors.New("'Name' is required on register")
		}
	} else {
		c.Name = "dummy"
	}

	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
