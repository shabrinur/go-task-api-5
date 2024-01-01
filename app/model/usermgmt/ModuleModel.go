package model

import (
	"time"

	"gorm.io/gorm"
)

type ModuleModel struct {
	ID          uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	Name        string         `json:"name" gorm:"type:varchar(100);required;unique;not null"`
	Type        string         `json:"type" gorm:"type:varchar(20);default:user"`
	Path        string         `json:"path" gorm:"type:varchar(200);required;unique;not null"`
	Description string         `json:"description" gorm:"type:varchar(100)"`
}

func (c *ModuleModel) TableName() string {
	return "app_module"
}
