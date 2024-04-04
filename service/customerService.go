package service

import (
	"github.com/scarecrow-404/banking/domain"
	"github.com/scarecrow-404/banking/errs"
)
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string)  (*domain.Customer, *errs.AppError)
	//GetCustomerByStatus(string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active"{
		status = "1"
	}else if status == "inactive"{
		status = "0"
	}else{
		status = ""
	}
	return s.repo.FindAll(status )
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)		
}

// func (s DefaultCustomerService) GetCustomerByStatus(status string) ([]domain.Customer, *errs.AppError) {
// 	return s.repo.FindStatus(status)
// }

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}