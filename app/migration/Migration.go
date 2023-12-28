package migration

import (
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/migration/masterdata"
	"idstar-idp/rest-api/app/util"
)

func Exec(pwdUtil util.PasswordUtil) error {
	db := config.GetUserMgmtDB()

	populateRole := masterdata.PopulateRole{}
	if err := populateRole.Exec(db); err != nil {
		return err
	}

	populateModule := masterdata.PopulateModule{}
	if err := populateModule.Exec(db); err != nil {
		return err
	}

	populateRoleModule := masterdata.PopulateRoleModule{}
	if err := populateRoleModule.Exec(db); err != nil {
		return err
	}

	pwd, err := pwdUtil.Encrypt("superadmin123")
	if err != nil {
		return err
	}
	createSuperAdmin := masterdata.CreateSuperAdmin{}
	if err := createSuperAdmin.Exec(db, *pwd); err != nil {
		return err
	}

	pwd, err = pwdUtil.Encrypt("v1admin123")
	if err != nil {
		return err
	}
	createV1Admin := masterdata.CreateV1Admin{}
	if err := createV1Admin.Exec(db, *pwd); err != nil {
		return err
	}

	return nil
}
