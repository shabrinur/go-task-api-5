package request

import (
	"idstar-idp/rest-api/app/util"

	"github.com/go-playground/validator/v10"
)

type TrainingRequest struct {
	Id       uint   `json:"id,omitempty" validate:"numeric"`
	Tema     string `json:"tema" validate:"required"`
	Pengajar string `json:"pengajar" validate:"required"`
}

func (c *TrainingRequest) Validate(isUpdate bool) error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	if isUpdate {
		if err := util.ValidateUpdateId(c.Id); err != nil {
			return err
		}
	}

	return nil
}
