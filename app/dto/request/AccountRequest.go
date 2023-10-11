package request

import (
	"idstar-idp/rest-api/app/util"

	"github.com/go-playground/validator/v10"
)

type AccountRequest struct {
	Id       uint      `json:"id,omitempty" validate:"numeric"`
	Nama     string    `json:"nama" validate:"required"`
	Jenis    string    `json:"jenis" validate:"required"`
	Rekening string    `json:"rekening" validate:"required"`
	Karyawan IdRequest `json:"karyawan"`
}

func (c *AccountRequest) Validate(isUpdate bool) error {
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
