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

func (repo *UserRepository) CheckUserAlreadyExists(username string) bool {
	var count int64
	repo.db.Model([]*model.RoleModel{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (repo *UserRepository) Create(dbObj *model.UserModel) (*model.UserModel, error) {
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error create new user, reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *UserRepository) UpdateOtp(dbObj *model.UserModel) (*model.UserModel, error) {
	result := repo.db.Model(&model.UserModel{}).Where("username = ?", dbObj.Username).Updates(model.UserModel{Otp: dbObj.Otp, OtpExpiredDate: dbObj.OtpExpiredDate})
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update OTP for user: ", dbObj.Username, ", reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *UserRepository) UpdatePassword(dbObj *model.UserModel) (*model.UserModel, error) {
	result := repo.db.Model(&model.UserModel{}).Where("username = ?", dbObj.Username).Update("password", dbObj.Password)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update password for user: ", dbObj.Username, ", reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *UserRepository) EnableOauthLogin(dbObj *model.UserModel) (*model.UserModel, error) {
	result := repo.db.Model(&model.UserModel{}).Where("username = ?", dbObj.Username).Updates(model.UserModel{
		Fullname:                dbObj.Fullname,
		Oauth:                   dbObj.Oauth,
		OauthProvider:           dbObj.OauthProvider,
		AccessToken:             dbObj.AccessToken,
		AccessTokenExpiredDate:  dbObj.AccessTokenExpiredDate,
		RefreshToken:            dbObj.RefreshToken,
		RefreshTokenExpiredDate: dbObj.RefreshTokenExpiredDate})
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update oauth info for user: ", dbObj.Username, ", reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *UserRepository) UpdateOauthInfo(dbObj *model.UserModel) (*model.UserModel, error) {
	result := repo.db.Model(&model.UserModel{}).Where("username = ?", dbObj.Username).Updates(model.UserModel{
		AccessToken:             dbObj.AccessToken,
		AccessTokenExpiredDate:  dbObj.AccessTokenExpiredDate,
		RefreshToken:            dbObj.RefreshToken,
		RefreshTokenExpiredDate: dbObj.RefreshTokenExpiredDate})
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update oauth info for user: ", dbObj.Username, ", reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *UserRepository) ActivateUser(dbObj *model.UserModel) (*model.UserModel, error) {
	result := repo.db.Model(&model.UserModel{}).Where("username = ?", dbObj.Username).Update("account_activated", dbObj.AccountActivated)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error activate user: ", dbObj.Username, ", reason: ", result.Error.Error()))
	}
	return dbObj, nil
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
