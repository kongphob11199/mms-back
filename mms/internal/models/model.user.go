package models

import "time"

func (ModelUser) TableName() string {
	return "users"
}

type ModelUser struct {
	UserId     uint32     `json:"user_id" gorm:"primary_key"`
	Firstname  string     `json:"firstname" gorm:"type:varchar; not null"`
	Lastname   string     `json:"lastname" gorm:"type:varchar; not null"`
	Gender     Gender     `json:"gender" gorm:"type:varchar; not null"`
	Role       Role       `json:"role" gorm:"type:varchar; not null"`
	Birthday   time.Time  `json:"birthday" gorm:"type:date;not null"`
	CreateAt   time.Time  `json:"create_at" gorm:"type:date;not null"`
	CreateBy   string     `json:"create_by" gorm:"type:varchar; not null"`
	UpdateAt   time.Time  `json:"update_at" gorm:"type:date;not null"`
	UpdateBy   string     `json:"update_by" gorm:"type:varchar; not null"`
	StatusUser StatusUser `json:"status_user" gorm:"type:varchar; not null"`
	Username   string     `json:"username" gorm:"type:varchar; not null"`
	Password   string     `json:"password" gorm:"type:text; not null"`
	Token      string     `json:"token" gorm:"type:text;"`
}

type Gender string

const (
	UNKNOWN Gender = "UNKNOWN"
	MALE    Gender = "MALE"
	FEMALE  Gender = "FEMALE"
)

type Role string

const (
	CUSTOMER   Role = "CUSTOMER"
	EMPLOYEE   Role = "EMPLOYEE"
	ADMIN      Role = "ADMIN"
	SUPERADMIN Role = "SUPERADMIN"
)

type StatusUser string

const (
	ACTIVE   StatusUser = "ACTIVE"
	INACTIVE StatusUser = "INACTIVE"
	DELETE   StatusUser = "DELETE"
)
