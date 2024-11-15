package service

import "mms/internal/repository"

type Service struct {
	User UserService
}

type Deps struct {
	Repository *repository.Repositories
}

func NewService(deps Deps) *Service {
	return &Service{
		User: NewServiceUser(deps.Repository.User),
	}
}
