package repository

import "gorm.io/gorm"

type Repositories struct {
	User UserRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User: NewRepositoryUser(db),
	}
}
