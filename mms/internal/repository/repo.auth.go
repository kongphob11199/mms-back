package repository

import (
	"mms/internal/dto"

	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) Login(req *dto.AuthLoginReq) (*dto.AuthLoginResp, error) {
	authLogin := &dto.AuthLoginResp{
		Token:    "",
		Response: "OK",
	}
	return authLogin, nil
}

func (r *repositoryAuth) CheckAuth() (*dto.StatusResp, error) {
	res := &dto.StatusResp{
		Response: "OK",
	}
	return res, nil
}
