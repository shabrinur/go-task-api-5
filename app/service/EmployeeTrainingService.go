package service

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/repository"
)

type EmployeeTrainingService struct {
	repo repository.EmployeeTrainingRepository
}

func NewEmployeeTrainingService(repo repository.EmployeeTrainingRepository) *EmployeeTrainingService {
	return &EmployeeTrainingService{repo}
}

func (svc *EmployeeTrainingService) CreateEmployeeTraining(req request.EmployeeTrainingRequest) (*model.EmployeeTrainingModel, error) {
	dbObj := model.EmployeeTrainingModel{
		Tanggal:    req.TanggalParsed,
		IdKaryawan: req.Karyawan.Id,
		IdTraining: req.Training.Id,
	}

	return svc.repo.Create(&dbObj)
}

func (svc *EmployeeTrainingService) UpdateEmployeeTraining(req request.EmployeeTrainingRequest) (*model.EmployeeTrainingModel, error) {
	dbObj := model.EmployeeTrainingModel{
		Tanggal:    req.TanggalParsed,
		IdKaryawan: req.Karyawan.Id,
		IdTraining: req.Training.Id,
	}

	return svc.repo.Update(req.Id, &dbObj)
}

func (svc *EmployeeTrainingService) GetEmployeeTrainingById(id int) (*model.EmployeeTrainingModel, error) {
	return svc.repo.GetById(uint(id))
}

func (svc *EmployeeTrainingService) GetEmployeeTrainingList(req request.PagingRequest) (*response.PaginationData, error) {
	pagedData := response.PaginationData{
		Pageable: req.Pageable,
		Sortable: req.Sortable,
	}
	return svc.repo.GetList(&pagedData)
}

func (svc *EmployeeTrainingService) DeleteEmployeeTraining(req request.IdRequest) error {
	return svc.repo.Delete(req.Id)
}
