package service

import (
	"github.com/arstrel/rest-banking/domain"
	"github.com/arstrel/rest-banking/errs"
)

// Interactors are classes that orchestrate and perform domain actions —
// think of Service Objects or Use Case Objects.
// They implement complex business rules and validation logic
// specific to a domain action (e.g., onboarding a production)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
