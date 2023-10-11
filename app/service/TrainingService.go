package service

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/repository"
)

type TrainingService struct {
	repo repository.TrainingRepository
}

func NewTrainingService(repo repository.TrainingRepository) *TrainingService {
	return &TrainingService{repo}
}

func (svc *TrainingService) CreateTraining(req request.TrainingRequest) (*model.TrainingModel, error) {
	dbObj := model.TrainingModel{
		Pengajar: req.Pengajar,
		Tema:     req.Tema,
	}

	return svc.repo.Create(&dbObj)
}

func (svc *TrainingService) UpdateTraining(req request.TrainingRequest) (*model.TrainingModel, error) {
	dbObj := model.TrainingModel{
		Pengajar: req.Pengajar,
		Tema:     req.Tema,
	}

	return svc.repo.Update(req.Id, &dbObj)
}

func (svc *TrainingService) GetTrainingById(id int) (*model.TrainingModel, error) {
	return svc.repo.GetById(uint(id))
}

func (svc *TrainingService) GetTrainingList(req request.PagingRequest) (*response.PaginationData, error) {
	pagedData := response.PaginationData{
		Pageable: req.Pageable,
		Sortable: req.Sortable,
	}
	return svc.repo.GetList(&pagedData)
}

func (svc *TrainingService) DeleteTraining(req request.IdRequest) error {
	return svc.repo.Delete(req.Id)
}
