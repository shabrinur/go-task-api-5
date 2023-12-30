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

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo}
}

func (svc *EmployeeService) CreateEmployee(req request.EmployeeRequest) (*model.EmployeeModel, int, error) {
	err := req.Validate(false)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = req.ParseDob()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	detail := model.EmployeeDetailModel{
		Nik:  req.DetailKaryawan.Nik,
		Npwp: req.DetailKaryawan.Npwp,
	}
	createdDetail, err := svc.repo.CreateDetail(&detail)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	dbObj := model.EmployeeModel{
		ID:       createdDetail.ID,
		Nama:     req.Nama,
		Dob:      req.DobParsed,
		Status:   req.Status,
		Alamat:   req.Alamat,
		IdDetail: createdDetail.ID,
	}
	result, err := svc.repo.Create(&dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, 0, nil
}

func (svc *EmployeeService) UpdateEmployee(req request.EmployeeRequest) (*model.EmployeeModel, int, error) {
	err := req.Validate(true)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = req.ParseDob()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	detail := model.EmployeeDetailModel{
		Nik:  req.DetailKaryawan.Nik,
		Npwp: req.DetailKaryawan.Npwp,
	}
	updatedDetail, err := svc.repo.UpdateDetail(req.DetailKaryawan.Id, &detail)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if updatedDetail == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for detail karyawan ID: ", req.DetailKaryawan.Id))
	}

	dbObj := model.EmployeeModel{
		Nama:     req.Nama,
		Dob:      req.DobParsed,
		Status:   req.Status,
		Alamat:   req.Alamat,
		IdDetail: updatedDetail.ID,
	}
	result, err := svc.repo.Update(req.Id, &dbObj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for karyawan ID: ", req.Id))
	}

	return result, 0, nil
}

func (svc *EmployeeService) GetEmployeeById(id int) (*model.EmployeeModel, int, error) {
	result, err := svc.repo.GetById(uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result == nil {
		return nil, http.StatusNotFound, errors.New(fmt.Sprint("record not exist for karyawan ID: ", id))
	}

	return result, 0, nil
}

func (svc *EmployeeService) GetEmployeeList(req request.PagingRequest) (*rsdata.PaginationData, int, error) {
	validFields := []string{"id", "nama", "status", "dob", "detail_karyawan", "created_date", "updated_date"}
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

func (svc *EmployeeService) DeleteEmployee(req request.IdRequest) (int, error) {
	err := req.Validate()
	if err != nil {
		return http.StatusBadRequest, err
	}

	deleted, err := svc.repo.Delete(req.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if deleted <= 0 {
		return http.StatusNotFound, errors.New(fmt.Sprint("record not exist for karyawan ID: ", req.Id))
	}

	return 0, nil
}
