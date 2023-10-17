package request

import (
	"idstar-idp/rest-api/app/util"
	"time"

	"github.com/go-playground/validator/v10"
)

type EmployeeTrainingRequest struct {
	Id            uint      `json:"id,omitempty" validate:"numeric"`
	Karyawan      IdRequest `json:"karyawan"`
	Training      IdRequest `json:"training"`
	Tanggal       string    `json:"tanggal" validate:"required,datetime=2006-01-02 15:04:05"`
	TanggalParsed time.Time `json:"-"`
}

func (c *EmployeeTrainingRequest) Validate(isUpdate bool) error {
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

func (c *EmployeeTrainingRequest) ParseTanggal() error {
	layout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, c.Tanggal)
	if err != nil {
		return err
	}
	c.TanggalParsed = date
	return nil
}
