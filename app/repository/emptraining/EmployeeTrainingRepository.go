package repository

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto/response"
	model "idstar-idp/rest-api/app/model/emptraining"
	"idstar-idp/rest-api/app/util"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeTrainingRepository struct {
	db *gorm.DB
}

func NewEmployeeTrainingRepository() *EmployeeTrainingRepository {
	return &EmployeeTrainingRepository{
		db: config.GetTrainingDB(),
	}
}

func (repo *EmployeeTrainingRepository) CheckDependencyRecordExists(idKaryawan uint, idTraining uint) error {
	if !util.CheckRecordExists([]*model.EmployeeModel{}, idKaryawan, repo.db) {
		return errors.New(fmt.Sprint("record not exist for karyawan ID: ", idKaryawan))
	}

	if !util.CheckRecordExists([]*model.TrainingModel{}, idTraining, repo.db) {
		return errors.New(fmt.Sprint("record not exist for training ID: ", idTraining))
	}
	return nil
}

func (repo *EmployeeTrainingRepository) Create(dbObj *model.EmployeeTrainingModel) (*model.EmployeeTrainingModel, error) {
	result := repo.db.Preload("Karyawan.DetailKaryawan").Preload(clause.Associations).Create(dbObj).First(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error create new karyawan training, reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *EmployeeTrainingRepository) Update(id uint, dbObj *model.EmployeeTrainingModel) (*model.EmployeeTrainingModel, error) {
	result := repo.db.Where("id = ?", id).Preload("Karyawan.DetailKaryawan").Preload(clause.Associations).Updates(dbObj).First(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update karyawan training with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *EmployeeTrainingRepository) GetById(id uint) (*model.EmployeeTrainingModel, error) {
	dbObj := &model.EmployeeTrainingModel{}
	result := repo.db.Where("id = ?", id).Preload("Karyawan.DetailKaryawan").Preload(clause.Associations).Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error find karyawan training with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}
func (repo *EmployeeTrainingRepository) GetList(pagedData *response.PaginationData) (*response.PaginationData, error) {
	dbObjs := []*model.EmployeeTrainingModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Preload("Karyawan.DetailKaryawan").Preload(clause.Associations).Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New(fmt.Sprint("error get karyawan training list, reason: ", result.Error.Error()))
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *EmployeeTrainingRepository) Delete(id uint) (int64, error) {
	result := repo.db.Where("id = ?", id).Delete(&model.EmployeeTrainingModel{})
	if result.Error != nil {
		return 0, errors.New(fmt.Sprint("error delete karyawan training with ID: ", id, ", reason: ", result.Error.Error()))
	}
	return result.RowsAffected, nil
}
