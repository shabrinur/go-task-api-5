package service

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response"
	model "idstar-idp/rest-api/app/model/emptraining"
	repository "idstar-idp/rest-api/app/repository/emptraining"
	"net/http"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{repo}
}

func (svc *AccountService) CreateAccount(req request.AccountRequest) (*model.AccountModel, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = svc.repo.CheckDependencyRecordExists(req.Karyawan.Id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	dbObj := model.AccountModel{
		Nama:       req.Nama,
		Jenis:      req.Jenis,
		Rekening:   req.Rekening,
		IDKaryawan: req.Karyawan.Id,
	}
	result, err := svc.repo.Create(&dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, 0, nil
}

func (svc *AccountService) UpdateAccount(req request.AccountRequest) (*model.AccountModel, int, error) {
	err := req.Validate(true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = svc.repo.CheckDependencyRecordExists(req.Karyawan.Id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	dbObj := model.AccountModel{
		Nama:       req.Nama,
		Jenis:      req.Jenis,
		Rekening:   req.Rekening,
		IDKaryawan: req.Karyawan.Id,
	}
	result, err := svc.repo.Update(req.Id, &dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for rekening ID: ", req.Id))
	}

	return result, 0, nil
}

func (svc *AccountService) GetAccountById(id int) (*model.AccountModel, int, error) {
	result, err := svc.repo.GetById(uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for rekening ID: ", id))
	}

	return result, 0, nil
}

func (svc *AccountService) GetAccountList(req request.PagingRequest) (*response.PaginationData, int, error) {
	validFields := []string{"id", "nama", "jenis", "rekening", "id_karyawan", "created_date", "updated_date"}
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

func (svc *AccountService) DeleteAccount(req request.IdRequest) (int, error) {
	err := req.Validate()
	if err != nil {
		return http.StatusBadRequest, err
	}

	deleted, err := svc.repo.Delete(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if deleted <= 0 {
		return http.StatusNotFound, errors.New(fmt.Sprint("record not exist for rekening ID: ", req.Id))
	}

	return 0, nil
}
