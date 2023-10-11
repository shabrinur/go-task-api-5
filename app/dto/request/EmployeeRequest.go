package request

import (
	"errors"
	"idstar-idp/rest-api/app/util"
	"time"

	"github.com/go-playground/validator/v10"
)

type EmployeeRequest struct {
	Id             uint           `json:"id,omitempty" validate:"numeric"`
	Nama           string         `json:"nama" validate:"required"`
	Dob            string         `json:"dob" validate:"required,datetime=2006-01-02"`
	Status         string         `json:"status" validate:"required"`
	Alamat         string         `json:"alamat" validate:"required"`
	DetailKaryawan DetailKaryawan `json:"detailKaryawan"`
}

type DetailKaryawan struct {
	Id   uint   `json:"id,omitempty" validate:"numeric"`
	Nik  string `json:"nik" validate:"required"`
	Npwp string `json:"npwp" validate:"required"`
}

func (c *EmployeeRequest) Validate(isUpdate bool) error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	if isUpdate {
		if err := util.ValidateUpdateId(c.Id); err != nil {
			return err
		}
		if c.DetailKaryawan.Id != c.Id {
			return errors.New("ID detail not equals to ID karyawan")
		}
	}

	return nil
}

func (c *EmployeeRequest) GetDob() (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, c.Dob)
}
