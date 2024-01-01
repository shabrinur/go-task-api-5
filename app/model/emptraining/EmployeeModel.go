package model

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeModel struct {
	ID             uint                `json:"id" gorm:"primaryKey;type:uint"`
	CreatedDate    time.Time           `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate    *time.Time          `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate    gorm.DeletedAt      `json:"deleted_date" gorm:"softDelete:true"`
	Alamat         string              `json:"alamat"`
	Dob            time.Time           `json:"dob" gorm:"required;not null"`
	Nama           string              `json:"nama" gorm:"type:varchar(100);required;not null"`
	Status         string              `json:"status" gorm:"type:varchar(20);required;not null"`
	IdDetail       uint                `json:"-" gorm:"column:detail_karyawan"`
	DetailKaryawan EmployeeDetailModel `json:"detailKaryawan" gorm:"foreignKey:IdDetail;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (c *EmployeeModel) TableName() string {
	return "karyawan"
}

type AccountEmployeeModel struct {
	ID   uint   `json:"id" gorm:"primaryKey;type:uint"`
	Nama string `json:"nama" gorm:"type:varchar(100);required;not null"`
}

func (c *AccountEmployeeModel) TableName() string {
	return "karyawan"
}
