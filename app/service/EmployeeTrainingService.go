package service

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/repository"
	"net/http"
)

type EmployeeTrainingService struct {
	repo repository.EmployeeTrainingRepository
}

func NewEmployeeTrainingService(repo repository.EmployeeTrainingRepository) *EmployeeTrainingService {
	return &EmployeeTrainingService{repo}
}

func (svc *EmployeeTrainingService) CreateEmployeeTraining(req request.EmployeeTrainingRequest) (*model.EmployeeTrainingModel, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = req.ParseTanggal()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = svc.repo.CheckDependencyRecordExists(req.Karyawan.Id, req.Training.Id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	dbObj := model.EmployeeTrainingModel{
		Tanggal:    req.TanggalParsed,
		IdKaryawan: req.Karyawan.Id,
		IdTraining: req.Training.Id,
	}
	result, err := svc.repo.Create(&dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, 0, nil
}

func (svc *EmployeeTrainingService) UpdateEmployeeTraining(req request.EmployeeTrainingRequest) (*model.EmployeeTrainingModel, int, error) {
	err := req.Validate(true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = req.ParseTanggal()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = svc.repo.CheckDependencyRecordExists(req.Karyawan.Id, req.Training.Id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	dbObj := model.EmployeeTrainingModel{
		Tanggal:    req.TanggalParsed,
		IdKaryawan: req.Karyawan.Id,
		IdTraining: req.Training.Id,
	}
	result, err := svc.repo.Update(req.Id, &dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for karyawan training ID: ", req.Id))
	}

	return result, 0, nil
}

func (svc *EmployeeTrainingService) GetEmployeeTrainingById(id int) (*model.EmployeeTrainingModel, int, error) {
	result, err := svc.repo.GetById(uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for karyawan training ID: ", id))
	}

	return result, 0, nil
}

func (svc *EmployeeTrainingService) GetEmployeeTrainingList(req request.PagingRequest) (*response.PaginationData, int, error) {
	validFields := []string{"id", "tanggal", "id_karyawan", "id_training", "created_date", "updated_date"}
	err := req.Validate(validFields)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	pagedData := response.PaginationData{
		Pageable: req.Pageable,
		Sortable: req.Sortable,
	}
	result, err := svc.repo.GetList(&pagedData)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, 0, nil
}

func (svc *EmployeeTrainingService) DeleteEmployeeTraining(req request.IdRequest) (int, error) {
	err := req.Validate()
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = svc.repo.Delete(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}
