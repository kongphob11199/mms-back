package service

import (
	"time"

	"mms/internal/dto"
	"mms/internal/middleware"
	"mms/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type ServiceAuth struct {
	Repository repository.AuthRepository
	UserRepo   repository.UserRepository
}

func NewServiceAuth(repo repository.AuthRepository, userRepo repository.UserRepository) *ServiceAuth {
	return &ServiceAuth{Repository: repo, UserRepo: userRepo}
}

func (s *ServiceAuth) Login(req *dto.AuthLoginReq) (*dto.AuthLoginResp, error) {
	user, err := s.UserRepo.FindUserByUsername(req)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId,
		"exp":     time.Now().Add(time.Hour * 14).Unix(),
	})

	tokenString, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		// log.Println("failed to generate token: %v", err)
		return nil, err
	}

	newReq := &dto.AuthUpdateTokenReq{
		UserId: user.UserId,
		Token:  tokenString,
	}

	if _, err := s.Repository.Login(newReq); err != nil {
		return nil, err
	}

	newResp := &dto.AuthLoginResp{
		Token:    tokenString,
		Response: dto.OK,
	}

	return newResp, nil
}

func (s *ServiceAuth) CheckAuth(req *dto.AuthUpdateTokenReq) (*dto.StatusResp, error) {
	res, err := s.Repository.CheckAuth(req)
	return res, err
}

func (s *ServiceAuth) Logout(req *dto.AuthUpdateTokenReq) (*dto.StatusResp, error) {
	newReq := &dto.AuthUpdateTokenReq{
		UserId: req.UserId,
		Token:  req.Token,
	}

	res, err := s.Repository.Logout(newReq)
	return res, err
}
