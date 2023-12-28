package model

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeDetailModel struct {
	ID          uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	Nik         string         `json:"nik" gorm:"type:varchar(20);required;not null"`
	Npwp        string         `json:"npwp" gorm:"type:varchar(20);required;not null"`
}

func (c *EmployeeDetailModel) TableName() string {
	return "training.detail_karyawan"
}
