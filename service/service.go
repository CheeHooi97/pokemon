package service

import "template/repository"

type Services struct {
	UserService    *UserService
	AdminService   *AdminService
	CompanyService *CompanyService
	CardService    CardService
}

func InitializeService(repos *repository.Repositories) *Services {
	return &Services{
		UserService:    NewUserService(repos.UserRepo),
		AdminService:   NewAdminService(repos.AdminRepo),
		CompanyService: NewCompanyService(repos.CompanyRepo),
		CardService:    NewCardService(repos.CardRepo),
	}
}
