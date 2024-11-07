package service

import (
	"mms/internal/dto"
	"mms/internal/models"
)

type UserService interface {
	FindAll() (*[]models.ModelUser, int64, error)
	FindPagination(req *dto.UserPaginationReq) (*[]models.ModelUser, int64, error)
	FindById(userId uint32) (*models.ModelUser, error)
	CreateCustomer(req *dto.CreateUserCustomerReq) (*dto.StatusResp, error)
	UpdateCustomer(userId uint32, req *dto.UpdateUserCustomerReq) (*dto.StatusResp, error)
	Create(req *dto.CreateUserReq) (*dto.StatusResp, error)
	Update(userId uint32, req *dto.UpdateUserReq) (*dto.StatusResp, error)
	Delete(userId uint32) (*dto.StatusResp, error)
	UpdateStatus(req *dto.UserStatusReq) (*dto.StatusResp, error)
}
