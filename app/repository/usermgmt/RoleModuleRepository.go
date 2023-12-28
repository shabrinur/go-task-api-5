package repository

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto"

	"gorm.io/gorm"
)

type RoleModuleRepository struct {
	db *gorm.DB
}

func NewRoleModuleRepository() *RoleModuleRepository {
	return &RoleModuleRepository{
		db: config.GetUserMgmtDB(),
	}
}

func (repo *RoleModuleRepository) GetPermissions(roleID uint) ([]dto.Permission, error) {
	permissions := []dto.Permission{}

	result := repo.db.Table("usermanagement.role_module").
		Select("app_module.path, role_module.get_allowed, role_module.put_allowed, role_module.post_allowed, role_module.delete_allowed").
		Joins("left join usermanagement.app_module on app_module.id = role_module.id_module").Where("role_module.id_role = ?", roleID).Scan(&permissions)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error get user permissions, reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, errors.New("invalid role permission config")
	}
	return permissions, nil
}
