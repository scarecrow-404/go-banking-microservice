package service

import (
	"github.com/scarecrow-404/banking/domain"
	"github.com/scarecrow-404/banking/errs"
)
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string)  (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)		
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}