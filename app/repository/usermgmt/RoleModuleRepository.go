package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto"
	model "idstar-idp/rest-api/app/model/usermgmt"

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

func (repo *RoleModuleRepository) GetDefaultUserRole() (*model.RoleModel, error) {
	dbObj := &model.RoleModel{}
	result := repo.db.Where("role_type = ? AND is_default = ?", "user", true).First(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error find default role, reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		result = repo.db.Where("role_type = ?", "user").First(&dbObj)
		if result.Error != nil {
			return nil, errors.New(fmt.Sprint("error find default role, reason: ", result.Error.Error()))
		}
		dbObj.IsDefault = sql.NullBool{Valid: true, Bool: true}
		repo.db.Where("id = ?", dbObj.ID).Updates(dbObj)
		if result.Error != nil {
			return nil, errors.New(fmt.Sprint("error find default role, reason: ", result.Error.Error()))
		}
	}
	return dbObj, nil
}

func (repo *RoleModuleRepository) GetPermissions(roleID uint) ([]dto.Permission, error) {
	permissions := []dto.Permission{}

	roleModule := &model.RoleModuleModel{}
	module := &model.ModuleModel{}

	result := repo.db.Table(roleModule.TableName()).
		Select("app_module.path, role_module.get_allowed, role_module.put_allowed, role_module.post_allowed, role_module.delete_allowed").
		Joins("left join "+module.TableName()+" on app_module.id = role_module.id_module").Where("role_module.id_role = ?", roleID).Scan(&permissions)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error get user permissions, reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, errors.New("invalid role permission config")
	}
	return permissions, nil
}
