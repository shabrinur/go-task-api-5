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

type TrainingRepository struct {
	db *gorm.DB
}

func NewTrainingRepository() *TrainingRepository {
	return &TrainingRepository{
		db: config.GetDB(),
	}
}

func (repo *TrainingRepository) Create(dbObj *model.TrainingModel) (*model.TrainingModel, error) {
	result := repo.db.Create(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error create new training, reason: ", result.Error.Error()))
	}
	return dbObj, nil
}

func (repo *TrainingRepository) Update(id uint, dbObj *model.TrainingModel) (*model.TrainingModel, error) {
	result := repo.db.Where("id = ?", id).Updates(dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error update training with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *TrainingRepository) GetById(id uint) (*model.TrainingModel, error) {
	dbObj := &model.TrainingModel{}
	result := repo.db.Where("id = ?", id).Find(&dbObj)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprint("error find training with ID: ", id, ", reason: ", result.Error.Error()))
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}
	return dbObj, nil
}

func (repo *TrainingRepository) GetList(pagedData *response.PaginationData) (*response.PaginationData, error) {
	dbObjs := []*model.TrainingModel{}
	util.CountRowsAndPages(dbObjs, pagedData, repo.db)

	if pagedData.TotalElements > 0 {
		result := repo.db.Scopes(util.Paginate(pagedData, repo.db)).Find(&dbObjs)
		if result.Error != nil {
			return nil, errors.New(fmt.Sprint("error get training list, reason: ", result.Error.Error()))
		}
		pagedData.Content = &dbObjs
		pagedData.NumberOfElements = int(result.RowsAffected)
	} else {
		pagedData.NumberOfElements = 0
	}
	pagedData.SetValueBeforeReturn()
	return pagedData, nil
}

func (repo *TrainingRepository) Delete(id uint) error {
	result := repo.db.Where("id_training = ?", id).Delete(&model.EmployeeTrainingModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("error delete karyawan training with training ID: ", id, ", reason: ", result.Error.Error()))
	}

	result = repo.db.Where("id = ?", id).Delete(&model.TrainingModel{})
	if result.Error != nil {
		return errors.New(fmt.Sprint("error delete training with ID: ", id, ", reason: ", result.Error.Error()))
	}
	return nil
}
