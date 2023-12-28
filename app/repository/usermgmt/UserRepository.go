package repository

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	model "idstar-idp/rest-api/app/model/usermgmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: config.GetUserMgmtDB(),
	}
}

func (repo *UserRepository) GetByUsername(username string) (*model.UserModel, error) {
	dbObj := &model.UserModel{}
	result := repo.db.Where("username = ?", username).Preload(clause.Associations).Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error find user with username: ", username, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}
