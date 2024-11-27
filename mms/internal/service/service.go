package service

import "mms/internal/repository"

type Service struct {
	User UserService
	Auth AuthService
}

type Deps struct {
	Repository *repository.Repositories
}

func NewService(deps Deps) *Service {
	return &Service{
		User: NewServiceUser(deps.Repository.User),
		Auth: NewServiceAuth(deps.Repository.Auth,deps.Repository.User),
	}
}
