package model

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeTrainingModel struct {
	ID          uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	Tanggal     time.Time      `json:"tanggal" gorm:"required;type:timestamp without time zone;not null"`
	IdKaryawan  uint           `json:"-"`
	Karyawan    EmployeeModel  `json:"karyawan" gorm:"foreignKey:IdKaryawan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IdTraining  uint           `json:"-"`
	Training    TrainingModel  `json:"training" gorm:"foreignKey:IdTraining;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (c *EmployeeTrainingModel) TableName() string {
	return "training.karyawan_training"
}
