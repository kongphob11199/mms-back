package dto

import (
	"mms/internal/models"
	"time"
)

type UserPaginationReq struct {
	Page       uint32              `json:"page"`
	PageLimit  uint32              `json:"pageLimit"`
	Search     string              `json:"search"`
	Role       []models.Role       `json:"role"`
	StatusUser []models.StatusUser `json:"statusUser"`
}

type CreateUserCustomerReq struct {
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Gender    models.Gender `json:"gender"`
	Birthday  time.Time     `json:"birthday"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
}

type UpdateUserCustomerReq struct {
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Gender    models.Gender `json:"gender"`
	Birthday  time.Time     `json:"birthday"`
	Username  string        `json:"username"`
}

type CreateUserReq struct {
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Gender    models.Gender `json:"gender"`
	Role      models.Role   `json:"role"`
	Birthday  time.Time     `json:"birthday"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
}

type UpdateUserReq struct {
	Firstname  string            `json:"firstname"`
	Lastname   string            `json:"lastname"`
	Gender     models.Gender     `json:"gender"`
	Role       models.Role       `json:"role"`
	Birthday   time.Time         `json:"birthday"`
	StatusUser models.StatusUser `json:"statusUser"`
	Username   string            `json:"username"`
}

type UserStatusReq struct {
	UserId     uint32            `json:"userId"`
	StatusUser models.StatusUser `json:"statusUser"`
}
