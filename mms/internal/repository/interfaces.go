package repository

import (
	"mms/internal/dto"
	"mms/internal/models"
)

type UserRepository interface {
	FindAll() (*[]models.ModelUser, int64, error)
	FindPagination(req *dto.UserPaginationReq) (*[]models.ModelUser, int64, error)
	FindById(userId uint32) (*models.ModelUser, error)
	CreateCustomer(req *dto.CreateUserCustomerReq) (*dto.StatusResp, error)
	UpdateCustomer(userId uint32, req *dto.UpdateUserCustomerReq) (*dto.StatusResp, error)
	Create(req *dto.CreateUserReq) (*dto.StatusResp, error)
	Update(userId uint32, req *dto.UpdateUserReq) (*dto.StatusResp, error)
	Delete(userId uint32) (*dto.StatusResp, error)
	UpdateStatus(req *dto.UserStatusReq) (*dto.StatusResp, error)
	FindUserByUsername(req *dto.AuthLoginReq) (*dto.UserFindUsernameRes, error)
}

type AuthRepository interface {
	Login(req *dto.AuthUpdateTokenReq) (*dto.StatusResp, error)
	CheckAuth(req *dto.AuthUpdateTokenReq) (*models.ModelUser, error)
	Logout(req *dto.AuthUpdateTokenReq) (*dto.StatusResp, error)
}
