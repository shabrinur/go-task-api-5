package model

import (
	"time"

	"gorm.io/gorm"
)

type TrainingModel struct {
	ID          uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	Pengajar    string         `json:"pengajar" gorm:"column:pengajar;type:varchar(100);required;not null"`
	Tema        string         `json:"tema" gorm:"column:tema;type:varchar(100);required;not null"`
}

func (c *TrainingModel) TableName() string {
	return "training"
}
