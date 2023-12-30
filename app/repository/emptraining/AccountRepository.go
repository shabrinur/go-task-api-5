package repository

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	model "idstar-idp/rest-api/app/model/emptraining"
	"idstar-idp/rest-api/app/util"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		db: config.GetTrainingDB(),
	}
}

func (repo *AccountRepository) CheckDependencyRecordExists(idKaryawan uint) error {
	if !util.CheckRecordExists([]*model.EmployeeModel{}, idKaryawan, repo.db) {
		return errors.New(fmt.Sprint("record not exist for karyawan ID: ", idKaryawan))
	}
	return nil
}

func (repo *AccountRepository) Create(dbObj *model.AccountModel) (*model.AccountModel, error) {
	result := repo.db.Preload(clause.Associations).Create(dbObj).First(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error create new rekening, reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *AccountRepository) Update(id uint, dbObj *model.AccountModel) (*model.AccountModel, error) {
	result := repo.db.Where("id = ?", id).Preload(clause.Associations).Updates(dbObj).First(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update rekening with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *AccountRepository) GetById(id uint) (*model.AccountModel, error) {
	dbObj := &model.AccountModel{}
	result := repo.db.Where("id = ?", id).Preload(clause.Associations).Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error find rekening with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *AccountRepository) GetList(pagedData *rsdata.PaginationData) (*rsdata.PaginationData, error) {
	dbObjs := []*model.AccountModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Preload(clause.Associations).Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New(fmt.Sprint("error get rekening list, reason: ", result.Error.Error()))
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *AccountRepository) Delete(id uint) (int64, error) {
	result := repo.db.Where("id = ?", id).Delete(&model.AccountModel{})
	if result.Error != nil {
		return 0, errors.New(fmt.Sprint("error delete rekening with ID: ", id, ", reason: ", result.Error.Error()))
	}
	return result.RowsAffected, nil
}
