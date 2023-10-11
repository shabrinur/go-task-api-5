package request

import "github.com/go-playground/validator/v10"

type IdRequest struct {
	Id uint `json:"id" validate:"required,numeric"`
}

func (c *IdRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}
