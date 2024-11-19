package service

import (
	"mms/internal/dto"
	"mms/internal/repository"
)

type ServiceAuth struct {
	Repostiory repository.AuthRepository
}

func NewServiceAuth(repo repository.AuthRepository) *ServiceAuth {
	return &ServiceAuth{Repostiory: repo}
}

func (s *ServiceAuth) Login(req *dto.AuthLoginReq) (*dto.AuthLoginResp, error) {
	res, err := s.Repostiory.Login(req)
	return res, err
}

func (s *ServiceAuth) CheckAuth() (*dto.StatusResp, error) {
	res, err := s.Repostiory.CheckAuth()
	return res, err
}
