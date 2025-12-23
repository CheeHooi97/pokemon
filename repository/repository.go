package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepo    UserRepository
	AdminRepo   AdminRepository
	CompanyRepo CompanyRepository
	CardRepo    CardRepository
}

func InitializeRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:    NewUserRepository(db),
		AdminRepo:   NewAdminRepository(db),
		CompanyRepo: NewCompanyRepository(db),
		CardRepo:    NewCardRepository(db),
	}
}
