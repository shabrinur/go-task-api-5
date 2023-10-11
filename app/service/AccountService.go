package service

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/repository"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repo}
}

func (svc *AccountService) CreateAccount(req request.AccountRequest) (*model.AccountModel, error) {
	dbObj := model.AccountModel{
		Nama:       req.Nama,
		Jenis:      req.Jenis,
		Rekening:   req.Rekening,
		IDKaryawan: req.Karyawan.Id,
	}

	return svc.repo.Create(&dbObj)
}

func (svc *AccountService) UpdateAccount(req request.AccountRequest) (*model.AccountModel, error) {
	dbObj := model.AccountModel{
		Nama:       req.Nama,
		Jenis:      req.Jenis,
		Rekening:   req.Rekening,
		IDKaryawan: req.Karyawan.Id,
	}

	return svc.repo.Update(req.Id, &dbObj)
}

func (svc *AccountService) GetAccountById(id int) (*model.AccountModel, error) {
	return svc.repo.GetById(uint(id))
}

func (svc *AccountService) GetAccountList(req request.PagingRequest) (*response.PaginationData, error) {
	pagedData := response.PaginationData{
		Pageable: req.Pageable,
		Sortable: req.Sortable,
	}
	return svc.repo.GetList(&pagedData)
}

func (svc *AccountService) DeleteAccount(req request.IdRequest) error {
	return svc.repo.Delete(req.Id)
}
