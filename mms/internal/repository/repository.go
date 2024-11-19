package repository

import "gorm.io/gorm"

type Repositories struct {
	User UserRepository
	Auth AuthRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User: NewRepositoryUser(db),
		Auth: NewRepositoryAuth(db),
	}
}
