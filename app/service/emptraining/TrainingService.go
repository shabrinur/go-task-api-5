package service

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response/rsdata"
	model "idstar-idp/rest-api/app/model/emptraining"
	repository "idstar-idp/rest-api/app/repository/emptraining"
	"net/http"
)

type TrainingService struct {
	repo repository.TrainingRepository
}

func NewTrainingService(repo repository.TrainingRepository) *TrainingService {
	return &TrainingService{repo}
}

func (svc *TrainingService) CreateTraining(req request.TrainingRequest) (*model.TrainingModel, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	dbObj := model.TrainingModel{
		Pengajar: req.Pengajar,
		Tema:     req.Tema,
	}
	result, err := svc.repo.Create(&dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, 0, nil
}

func (svc *TrainingService) UpdateTraining(req request.TrainingRequest) (*model.TrainingModel, int, error) {
	err := req.Validate(true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	dbObj := model.TrainingModel{
		Pengajar: req.Pengajar,
		Tema:     req.Tema,
	}
	result, err := svc.repo.Update(req.Id, &dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for training ID: ", req.Id))
	}

	return result, 0, nil
}

func (svc *TrainingService) GetTrainingById(id int) (*model.TrainingModel, int, error) {
	result, err := svc.repo.GetById(uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for training ID: ", id))
	}

	return result, 0, nil
}

func (svc *TrainingService) GetTrainingList(req request.PagingRequest) (*rsdata.PaginationData, int, error) {
	validFields := []string{"id", "tema", "pengajar", "created_date", "updated_date"}
	err := req.Validate(validFields)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	pagedData := rsdata.PaginationData{
		Pageable: req.Pageable,
		Sortable: req.Sortable,
	}
	result, err := svc.repo.GetList(&pagedData)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, 0, nil
}

func (svc *TrainingService) DeleteTraining(req request.IdRequest) (int, error) {
	err := req.Validate()
	if err != nil {
		return http.StatusBadRequest, err
	}

	deleted, err := svc.repo.Delete(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if deleted <= 0 {
		return http.StatusNotFound, errors.New(fmt.Sprint("record not exist for training ID: ", req.Id))
	}

	return 0, nil
}
