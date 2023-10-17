package service

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/dto/response"
	"idstar-idp/rest-api/app/model"
	"idstar-idp/rest-api/app/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo}
}

func (svc *EmployeeService) CreateEmployee(req request.EmployeeRequest) (*model.EmployeeModel, error) {
	detail := model.EmployeeDetailModel{
		Nik:  req.DetailKaryawan.Nik,
		Npwp: req.DetailKaryawan.Npwp,
	}

	createdDetail, err := svc.repo.CreateDetail(&detail)
	if err != nil {
		return nil, err
	}

	dbObj := model.EmployeeModel{
		Nama:     req.Nama,
		Dob:      req.DobParsed,
		Status:   req.Status,
		Alamat:   req.Alamat,
		IdDetail: createdDetail.ID,
	}

	return svc.repo.Create(&dbObj)
}

func (svc *EmployeeService) UpdateEmployee(req request.EmployeeRequest) (*model.EmployeeModel, error) {
	detail := model.EmployeeDetailModel{
		Nik:  req.DetailKaryawan.Nik,
		Npwp: req.DetailKaryawan.Npwp,
	}

	createdDetail, err := svc.repo.UpdateDetail(req.DetailKaryawan.Id, &detail)
	if err != nil {
		return nil, err
	}

	dbObj := model.EmployeeModel{
		Nama:     req.Nama,
		Dob:      req.DobParsed,
		Status:   req.Status,
		Alamat:   req.Alamat,
		IdDetail: createdDetail.ID,
	}

	result, err := svc.repo.Create(&dbObj)
	if err != nil {
		return nil, err
	}
	result.DetailKaryawan = *createdDetail
	return result, nil
}

func (svc *EmployeeService) GetEmployeeById(id int) (*model.EmployeeModel, error) {
	return svc.repo.GetById(uint(id))
}

func (svc *EmployeeService) GetEmployeeList(req request.PagingRequest) (*response.PaginationData, error) {
	pagedData := response.PaginationData{
		Pageable: req.Pageable,
		Sortable: req.Sortable,
	}
	return svc.repo.GetList(&pagedData)
}

func (svc *EmployeeService) DeleteEmployee(req request.IdRequest) error {
	return svc.repo.Delete(req.Id)
}
