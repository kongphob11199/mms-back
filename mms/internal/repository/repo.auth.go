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
			Response: dto.ERROR,
		}, message.ErrorUserNotFound
	}

	user.Token = req.Token

	if err := r.db.Debug().Save(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: dto.ERROR,
		}, message.ErrorUserUPDATE
	}

	authLogin := &dto.StatusResp{
		Response: dto.OK,
	}
	return authLogin, nil
}

func (r *repositoryAuth) CheckAuth(req *dto.AuthUpdateTokenReq) (*models.ModelUser, error) {

	var user models.ModelUser

	if err := r.db.Where("user_id = ?", req.UserId).First(&user).Error; err != nil {
		return &models.ModelUser{}, message.ErrorUserNotFound
	}

	if user.Token != req.Token {
		return &models.ModelUser{}, message.ErrorInvalidToken
	}

	return &user, nil
}

func (r *repositoryAuth) Logout(req *dto.AuthUpdateTokenReq) (*dto.StatusResp, error) {

	var user models.ModelUser

	if err := r.db.Where("user_id = ?", req.UserId).First(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: dto.ERROR,
		}, message.ErrorUserNotFound
	}

	user.Token = req.Token

	if err := r.db.Debug().Save(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: dto.ERROR,
		}, message.ErrorUserUPDATE
	}

	res := &dto.StatusResp{
		Response: dto.OK,
	}
	return res, nil
}
