package repository

import (
	"mms/internal/dto"
	"mms/internal/message"
	"mms/internal/models"

	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) Login(req *dto.AuthUpdateTokenReq) (*dto.StatusResp, error) {

	var user models.ModelUser

	if err := r.db.Where("user_id = ?", req.UserId).First(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserNotFound
	}

	user.Token = req.Token

	authLogin := &dto.StatusResp{
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
