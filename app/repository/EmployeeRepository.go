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

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		db: config.GetDB(),
	}
}

func (repo *EmployeeRepository) CreateDetail(dbObj *model.EmployeeDetailModel) (*model.EmployeeDetailModel, error) {
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error create new detail karyawan, reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *EmployeeRepository) Create(dbObj *model.EmployeeModel) (*model.EmployeeModel, error) {
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error create new karyawan, reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *EmployeeRepository) UpdateDetail(id uint, dbObj *model.EmployeeDetailModel) (*model.EmployeeDetailModel, error) {
	result := repo.db.Where("id = ?", id).Updates(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update detail karyawan with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *EmployeeRepository) Update(id uint, dbObj *model.EmployeeModel) (*model.EmployeeModel, error) {
	result := repo.db.Where("id = ?", id).Updates(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update karyawan with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *EmployeeRepository) GetById(id uint) (*model.EmployeeModel, error) {
	dbObj := &model.EmployeeModel{}
	result := repo.db.Where("id = ?", id).Preload("DetailKaryawan").Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error find karyawan with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *EmployeeRepository) GetList(pagedData *response.PaginationData) (*response.PaginationData, error) {
	dbObjs := []*model.EmployeeModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Preload("DetailKaryawan").Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New(fmt.Sprint("error get karyawan list, reason: ", result.Error.Error()))
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *EmployeeRepository) Delete(id uint) error {
	result := repo.db.Where("id_karyawan = ?", id).Delete(&model.EmployeeTrainingModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("error delete karyawan training with karyawan ID: ", id, ", reason: ", result.Error.Error()))
	}

	result = repo.db.Where("id_karyawan = ?", id).Delete(&model.AccountModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("error delete rekening with karyawan ID: ", id, ", reason: ", result.Error.Error()))
	}

	result = repo.db.Where("detail_karyawan = ?", id).Delete(&model.EmployeeModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("error delete karyawan with ID: ", id, ", reason: ", result.Error.Error()))
	}

	result = repo.db.Where("id = ?", id).Delete(&model.EmployeeDetailModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("error delete detail karyawan with karyawan ID: ", id, ", reason: ", result.Error.Error()))
	}
	return nil
}
