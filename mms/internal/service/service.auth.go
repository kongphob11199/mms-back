package service

import (
	"log"
	"time"

	"mms/internal/middleware"
	"mms/internal/dto"
	"mms/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type ServiceAuth struct {
	Repository repository.AuthRepository
	UserRepo   repository.UserRepository
}

func NewServiceAuth(repo repository.AuthRepository, userRepo repository.UserRepository) *ServiceAuth {
	return &ServiceAuth{Repository: repo,  UserRepo: userRepo}
}

func (s *ServiceAuth) Login(req *dto.AuthLoginReq) (*dto.AuthLoginResp, error) {
	user,err := s.UserRepo.FindUserByUsername(req)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId,
		"exp":      time.Now().Add(time.Hour * 14).Unix(), 
	})

	tokenString, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		log.Println("failed to generate token: %v", err)
		return nil, err
	}

	newResp := &dto.AuthLoginResp{
		Token:tokenString,
		Response:dto.OK,
	}

	return newResp, nil
}

func (s *ServiceAuth) CheckAuth() (*dto.StatusResp, error) {
	res, err := s.Repository.CheckAuth()
	return res, err
}
