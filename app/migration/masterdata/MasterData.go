package masterdata

import (
	"database/sql"
	model "idstar-idp/rest-api/app/model/usermgmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GetDefaultRole() []model.RoleModel {
	return []model.RoleModel{
		{
			Name:      "superadmin",
			RoleType:  "superadmin",
			IsDefault: sql.NullBool{Valid: true, Bool: false},
		},
		{
			Name:      "v1-api admin",
			RoleType:  "admin",
			IsDefault: sql.NullBool{Valid: true, Bool: false},
		},
		{
			Name:      "idstar api user",
			RoleType:  "user",
			IsDefault: sql.NullBool{Valid: true, Bool: true},
		},
	}
}

func GetDefaultModule() []model.ModuleModel {
	return []model.ModuleModel{
		{
			Name:        "superadmin",
			Type:        "superadmin",
			Path:        "/",
			Description: "superadmin exclusive module",
		},
		{
			Name:        "v1 api",
			Type:        "admin",
			Path:        "/v1",
			Description: "all v1 api module",
		},
		{
			Name:        "idstar karyawan",
			Type:        "user",
			Path:        "/v1/idstar/karyawan",
			Description: "idstar karyawan module",
		},
		{
			Name:        "idstar rekening",
			Type:        "user",
			Path:        "/v1/idstar/rekening",
			Description: "idstar rekening module",
		},
		{
			Name:        "idstar training",
			Type:        "user",
			Path:        "/v1/idstar/training",
			Description: "idstar rekening module",
		},
		{
			Name:        "idstar karyawan-training",
			Type:        "user",
			Path:        "/v1/idstar/karyawan-training",
			Description: "idstar karyawan-training module",
		},
	}
}

func GetSuperAdmin() *model.UserModel {
	return &model.UserModel{
		Fullname:         "Super Admin",
		Username:         "superadmin@go-restapi.idp",
		Otp:              "DUMMYOTP",
		OtpExpiredDate:   time.Now(),
		AccountActivated: sql.NullBool{Valid: true, Bool: true},
	}
}

func GetV1Admin() *model.UserModel {
	return &model.UserModel{
		Fullname:         "v1 API Admin",
		Username:         "v1admin@go-restapi.idp",
		Otp:              "DUMMYOTP",
		OtpExpiredDate:   time.Now(),
		AccountActivated: sql.NullBool{Valid: true, Bool: true},
	}
}

type PopulateRole struct{}

func (d *PopulateRole) Exec(db *gorm.DB) error {
	model := model.RoleModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + model.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		roles := GetDefaultRole()
		for _, role := range roles {
			result := db.Create(&role)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

type PopulateModule struct{}

func (d *PopulateModule) Exec(db *gorm.DB) error {
	entity := model.ModuleModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + entity.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		roles := GetDefaultModule()
		for _, role := range roles {
			result := db.Create(&role)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

type PopulateRoleModule struct{}

func (d *PopulateRoleModule) Exec(db *gorm.DB) error {
	entity := model.RoleModuleModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + entity.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		roles := []*model.RoleModel{}
		err := db.Find(&roles).Error
		if err != nil {
			return err
		}
		modules := []*model.ModuleModel{}
		err = db.Find(&modules).Error
		if err != nil {
			return err
		}
		for _, role := range roles {
			for _, module := range modules {
				if strings.EqualFold(role.RoleType, module.Type) {
					roleModule := &model.RoleModuleModel{
						IdRole:        role.ID,
						IdModule:      module.ID,
						PostAllowed:   sql.NullBool{Valid: true, Bool: true},
						PutAllowed:    sql.NullBool{Valid: true, Bool: true},
						GetAllowed:    sql.NullBool{Valid: true, Bool: true},
						DeleteAllowed: sql.NullBool{Valid: true, Bool: true},
					}
					result := db.Create(&roleModule)
					if result.Error != nil {
						return result.Error
					}
				}
			}
		}
	}
	return nil
}

type CreateSuperAdmin struct{}

func (d *CreateSuperAdmin) Exec(db *gorm.DB, password string) error {
	role := &model.RoleModel{}
	result := db.Where("role_type = ?", "superadmin").First(&role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		entity := model.UserModel{}
		count := 0
		err := db.Raw("SELECT COUNT(*) FROM "+entity.TableName()+" WHERE id_role = ?", role.ID).Scan(&count).Error
		if err != nil {
			return err
		}

		if count == 0 {
			superAdmin := GetSuperAdmin()
			superAdmin.IDRole = role.ID
			superAdmin.Password = password
			result := db.Create(&superAdmin)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

type CreateV1Admin struct{}

func (d *CreateV1Admin) Exec(db *gorm.DB, password string) error {
	role := &model.RoleModel{}
	result := db.Where("role_type = ? AND name = ?", "admin", "v1-api admin").First(&role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		entity := model.UserModel{}
		count := 0
		err := db.Raw("SELECT COUNT(*) FROM "+entity.TableName()+" WHERE id_role = ?", role.ID).Scan(&count).Error
		if err != nil {
			return err
		}

		if count == 0 {
			v1Admin := GetV1Admin()
			v1Admin.IDRole = role.ID
			v1Admin.Password = password
			result := db.Create(&v1Admin)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}
