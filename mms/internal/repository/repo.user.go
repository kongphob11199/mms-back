package repository

import (
	"fmt"
	"mms/internal/dto"
	"mms/internal/message"
	"mms/internal/models"
	"mms/internal/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

func (r *repositoryUser) FindAll() (*[]models.ModelUser, int64, error) {

	var users []models.ModelUser
	var total int64

	err := r.db.Model(&models.ModelUser{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return &users, total, nil
}

func (r *repositoryUser) FindPagination(req *dto.UserPaginationReq) (*[]models.ModelUser, int64, error) {

	var users []models.ModelUser
	var total int64

	offset := int(req.Page * req.PageLimit)
	limit := int(req.PageLimit)

	query := r.db.Model(&models.ModelUser{})

	if req.Search != "" {
		query = query.Where("firstname LIKE ? OR lastname LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if len(req.Role) > 0 {
		query = query.Where("role IN ?", req.Role)
	}

	if len(req.StatusUser) > 0 {
		query = query.Where("status_user IN ?", req.StatusUser)
	}

	err := r.db.Model(&models.ModelUser{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return &users, total, nil
}

func (r *repositoryUser) FindById(userId uint32) (*models.ModelUser, error) {

	var users models.ModelUser

	query := r.db.Model(&models.ModelUser{})

	query = query.Where("user_id = ?", userId)

	err := query.Find(&users).Error
	if err != nil {
		return &models.ModelUser{}, err
	}

	return &users, nil
}

func (r *repositoryUser) CreateCustomer(req *dto.CreateUserCustomerReq) (*dto.StatusResp, error) {

	var user models.ModelUser

	if len(req.Username) < 4 {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUsernameLength
	}

	var existingUser models.ModelUser
	if err := r.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserDup
	}

	if err := utils.ValidatePassword(req.Password); !err {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorPassWordCheck
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorPasswordHash
	}

	user.Firstname = req.Firstname
	user.Lastname = req.Lastname
	user.Username = req.Username
	user.Password = string(hashedPassword)
	user.Gender = req.Gender
	user.Role = models.CUSTOMER
	user.Birthday = req.Birthday
	user.CreateAt = time.Now()
	user.CreateBy = fmt.Sprintf("%s %s", req.Firstname, req.Lastname)
	user.UpdateAt = time.Now()
	user.UpdateBy = fmt.Sprintf("%s %s", req.Firstname, req.Lastname)
	user.StatusUser = models.ACTIVE

	db := r.db.Model(&user)

	addUser := db.Debug().Create(&user).Commit()

	if addUser.RowsAffected < 1 {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserCreateCustomer
	}

	return &dto.StatusResp{
		Response: "OK",
	}, nil
}

func (r *repositoryUser) UpdateCustomer(userId uint32, req *dto.UpdateUserCustomerReq) (*dto.StatusResp, error) {

	var user models.ModelUser

	if len(req.Username) < 4 {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUsernameLength
	}

	if err := r.db.First(&user, userId).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserNotFound
	}

	err := r.db.Where("username = ? AND id != ?", req.Username, userId).First(&user).Error
	if err == nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserDup
	}

	user.Firstname = req.Firstname
	user.Lastname = req.Lastname
	user.Username = req.Username
	user.Gender = req.Gender
	user.Birthday = req.Birthday
	user.UpdateAt = time.Now()
	user.UpdateBy = fmt.Sprintf("%s %s", req.Firstname, req.Lastname)

	if err := r.db.Debug().Save(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserUpdateCustomer
	}

	return &dto.StatusResp{
		Response: "OK",
	}, nil
}

func (r *repositoryUser) Create(req *dto.CreateUserReq) (*dto.StatusResp, error) {

	var user models.ModelUser

	if len(req.Username) < 4 {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUsernameLength
	}

	var existingUser models.ModelUser
	if err := r.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserDup
	}

	if err := utils.ValidatePassword(req.Password); !err {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorPassWordCheck
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorPasswordHash
	}

	user.Firstname = req.Firstname
	user.Lastname = req.Lastname
	user.Username = req.Username
	user.Password = string(hashedPassword)
	user.Gender = req.Gender
	user.Role = req.Role
	user.Birthday = req.Birthday
	user.CreateAt = time.Now()
	user.CreateBy = "ADMIN"
	user.UpdateAt = time.Now()
	user.UpdateBy = "ADMIN"
	user.StatusUser = models.ACTIVE

	db := r.db.Model(&user)

	addUser := db.Debug().Create(&user).Commit()

	if addUser.RowsAffected < 1 {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserCreateCustomer
	}

	return &dto.StatusResp{
		Response: "OK",
	}, nil
}

func (r *repositoryUser) Update(userId uint32, req *dto.UpdateUserReq) (*dto.StatusResp, error) {

	var user models.ModelUser

	if len(req.Username) < 4 {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUsernameLength
	}

	if err := r.db.First(&user, userId).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserNotFound
	}

	err := r.db.Where("username = ? AND id != ?", req.Username, userId).First(&user).Error
	if err == nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserDup
	}

	user.Firstname = req.Firstname
	user.Lastname = req.Lastname
	user.Username = req.Username
	user.Gender = req.Gender
	user.Role = req.Role
	user.Birthday = req.Birthday
	user.UpdateAt = time.Now()
	user.UpdateBy = "ADMIN"
	user.StatusUser = req.StatusUser

	if err := r.db.Debug().Save(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserUpdateCustomer
	}

	return &dto.StatusResp{
		Response: "OK",
	}, nil
}

func (r *repositoryUser) Delete(userId uint32) (*dto.StatusResp, error) {

	var user models.ModelUser
	if err := r.db.First(&user, userId).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserNotFound
	}

	user.UpdateAt = time.Now()
	user.UpdateBy = "ADMIN"
	user.StatusUser = models.DELETE

	if err := r.db.Debug().Save(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserUpdateCustomer
	}

	return &dto.StatusResp{
		Response: "OK",
	}, nil
}

func (r *repositoryUser) UpdateStatus(req *dto.UserStatusReq) (*dto.StatusResp, error) {

	var user models.ModelUser
	if err := r.db.First(&user, req.UserId).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserNotFound
	}

	user.UpdateAt = time.Now()
	user.UpdateBy = "ADMIN"
	user.StatusUser = req.StatusUser

	if err := r.db.Debug().Save(&user).Error; err != nil {
		return &dto.StatusResp{
			Response: "ERROR",
		}, message.ErrorUserUpdateCustomer
	}

	return &dto.StatusResp{
		Response: "OK",
	}, nil
}

func (r *repositoryUser) FindUserByUsername (req *dto.AuthLoginReq) (*dto.UserFindUsernameRes, error) {

	var user models.ModelUser

	if err := r.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, message.ErrorUserNotFound_Login
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, message.ErrorInvalidPassword_Login
	}

	res := &dto.UserFindUsernameRes{
		UserId: user.UserId,
	}

	return res,nil
}
