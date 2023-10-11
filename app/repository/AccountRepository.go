package repository

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/util"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		db: config.GetDB(),
	}
}

func (repo *AccountRepository) Create(dbObj *model.AccountModel) (*model.AccountModel, error) {
	if !util.CheckRecordExists([]*model.EmployeeModel{}, dbObj.IDKaryawan, repo.db) {
		return nil, errors.New(fmt.Sprint("record not exist for karyawan ID: ", dbObj.IDKaryawan))
	}
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New("failed to create new rekening")
	}
	return dbObj, nil
}

func (repo *AccountRepository) Update(id uint, dbObj *model.AccountModel) (*model.AccountModel, error) {
	if !util.CheckRecordExists([]*model.EmployeeModel{}, dbObj.IDKaryawan, repo.db) {
		return nil, errors.New(fmt.Sprint("record not exist for karyawan ID: ", dbObj.IDKaryawan))

	}
	result := repo.db.Where("id = ?", id).Updates(dbObj)
	if result.Error != nil {
		return nil, errors.New("failed to create new rekening")
	}
	return dbObj, nil
}

func (repo *AccountRepository) GetById(id uint) (*model.AccountModel, error) {
	dbObj := &model.AccountModel{}
	result := repo.db.Where("id = ?", id).Preload("Karyawan").Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("failed to find rekening with ID: ", id))
	}
	return dbObj, nil
}

func (repo *AccountRepository) GetList(pagedData *response.PaginationData) (*response.PaginationData, error) {
	dbObjs := []*model.AccountModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Preload("Karyawan").Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New("failed to get rekening list")
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *AccountRepository) Delete(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&model.AccountModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("failed to delete rekening with ID: ", id))
	}
	return nil
}
