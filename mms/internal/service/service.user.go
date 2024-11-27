package service

import (
	"mms/internal/dto"
	"mms/internal/models"
	"mms/internal/repository"
)

type ServiceUser struct {
	Repository repository.UserRepository
}

func NewServiceUser(repo repository.UserRepository) *ServiceUser {
	return &ServiceUser{Repository: repo}
}

func (s *ServiceUser) FindAll() (*[]models.ModelUser, int64, error) {
	users, total, err := s.Repository.FindAll()

	return users, total, err
}

func (s *ServiceUser) FindPagination(req *dto.UserPaginationReq) (*[]models.ModelUser, int64, error) {
	users, total, err := s.Repository.FindPagination(req)

	return users, total, err
}

func (s *ServiceUser) FindById(userId uint32) (*models.ModelUser, error) {
	users, err := s.Repository.FindById(userId)

	return users, err
}

func (s *ServiceUser) CreateCustomer(req *dto.CreateUserCustomerReq) (*dto.StatusResp, error) {
	res, err := s.Repository.CreateCustomer(req)

	return res, err
}

func (s *ServiceUser) UpdateCustomer(userId uint32, req *dto.UpdateUserCustomerReq) (*dto.StatusResp, error) {
	res, err := s.Repository.UpdateCustomer(userId, req)

	return res, err
}

func (s *ServiceUser) Create(req *dto.CreateUserReq) (*dto.StatusResp, error) {
	res, err := s.Repository.Create(req)

	return res, err
}

func (s *ServiceUser) Update(userId uint32, req *dto.UpdateUserReq) (*dto.StatusResp, error) {
	res, err := s.Repository.Update(userId, req)

	return res, err
}

func (s *ServiceUser) Delete(userId uint32) (*dto.StatusResp, error) {
	res, err := s.Repository.Delete(userId)

	return res, err
}

func (s *ServiceUser) UpdateStatus(req *dto.UserStatusReq) (*dto.StatusResp, error) {
	res, err := s.Repository.UpdateStatus(req)

	return res, err
}
