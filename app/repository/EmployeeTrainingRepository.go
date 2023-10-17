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

type EmployeeTrainingRepository struct {
	db *gorm.DB
}

func NewEmployeeTrainingRepository() *EmployeeTrainingRepository {
	return &EmployeeTrainingRepository{
		db: config.GetDB(),
	}
}

func (repo *EmployeeTrainingRepository) Create(dbObj *model.EmployeeTrainingModel) (*model.EmployeeTrainingModel, error) {
	if !util.CheckRecordExists([]*model.EmployeeModel{}, dbObj.IdKaryawan, repo.db) {
		return nil, errors.New(fmt.Sprint("record not exist for karyawan ID: ", dbObj.IdKaryawan))
	}
	if !util.CheckRecordExists([]*model.TrainingModel{}, dbObj.IdTraining, repo.db) {
		return nil, errors.New(fmt.Sprint("record not exist for training ID: ", dbObj.IdTraining))
	}
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New("failed to create new karyawan training")
	}
	return dbObj, nil
}

func (repo *EmployeeTrainingRepository) Update(id uint, dbObj *model.EmployeeTrainingModel) (*model.EmployeeTrainingModel, error) {
	if !util.CheckRecordExists([]*model.EmployeeModel{}, dbObj.IdKaryawan, repo.db) {
		return nil, errors.New(fmt.Sprint("record not exist for karyawan ID: ", dbObj.IdKaryawan))
	}
	if !util.CheckRecordExists([]*model.TrainingModel{}, dbObj.IdTraining, repo.db) {
		return nil, errors.New(fmt.Sprint("record not exist for training ID: ", dbObj.IdTraining))
	}
	result := repo.db.Where("id = ?", id).Updates(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("failed to update karyawan training with ID: ", id))
	}
	return dbObj, nil
}

func (repo *EmployeeTrainingRepository) GetById(id uint) (*model.EmployeeTrainingModel, error) {
	dbObj := &model.EmployeeTrainingModel{}
	result := repo.db.Where("id = ?", id).Preload("Karyawan").Preload("Training").Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("failed to find karyawan training with ID: ", id))
	}
	return dbObj, nil
}
func (repo *EmployeeTrainingRepository) GetList(pagedData *response.PaginationData) (*response.PaginationData, error) {
	dbObjs := []*model.EmployeeTrainingModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Preload("Karyawan").Preload("Training").Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New("failed to get karyawan training list")
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *EmployeeTrainingRepository) Delete(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&model.EmployeeTrainingModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("failed to delete karyawan training with ID: ", id))
	}
	return nil
}
