package model

import (
	"time"

	"gorm.io/gorm"
)

type AccountModel struct {
	ID          uint                 `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate time.Time            `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate *time.Time           `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt       `json:"deleted_date" gorm:"softDelete:true"`
	Jenis       string               `json:"jenis" gorm:"type:varchar(20);required;not null"`
	Nama        string               `json:"nama" gorm:"type:varchar(100);required;not null"`
	Rekening    string               `json:"rekening" gorm:"type:varchar(20);required;not null"`
	IDKaryawan  uint                 `json:"-"`
	Karyawan    AccountEmployeeModel `json:"karyawan" gorm:"foreignKey:IDKaryawan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (c *AccountModel) TableName() string {
	return "rekening"
}
