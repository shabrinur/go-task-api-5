package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type RoleModel struct {
	ID          uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	Name        string         `json:"name" gorm:"type:varchar(50);required;unique;not null"`
	RoleType    string         `json:"roleType" gorm:"type:varchar(20);default:user"`
	IsDefault   sql.NullBool   `json:"-" gorm:"default:false"`
}

func (c *RoleModel) TableName() string {
	return "usermanagement.oauth_role"
}
