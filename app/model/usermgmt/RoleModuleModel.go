package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type RoleModuleModel struct {
	ID            uint           `json:"id" gorm:"primaryKey;type:uint;autoIncrement"`
	CreatedDate   time.Time      `json:"created_date" gorm:"autoCreateTime:true;not null"`
	UpdatedDate   *time.Time     `json:"updated_date" gorm:"autoUpdateTime:true"`
	DeletedDate   gorm.DeletedAt `json:"deleted_date" gorm:"softDelete:true"`
	GetAllowed    sql.NullBool   `json:"getAllowed" gorm:"default:false"`
	PutAllowed    sql.NullBool   `json:"putAllowed" gorm:"default:false"`
	PostAllowed   sql.NullBool   `json:"postAllowed" gorm:"default:false"`
	DeleteAllowed sql.NullBool   `json:"deleteAllowed" gorm:"default:false"`
	IdRole        uint           `json:"-"`
	Role          RoleModel      `json:"role" gorm:"foreignKey:IdRole;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IdModule      uint           `json:"-"`
	Module        ModuleModel    `json:"module" gorm:"foreignKey:IdModule;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (c *RoleModuleModel) TableName() string {
	return "role_module"
}
