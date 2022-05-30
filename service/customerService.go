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
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	statusToCodeMap := map[string]string{
		"active":   "1",
		"inactive": "0",
	}

	if sCode, ok := statusToCodeMap[status]; ok {
		status = sCode
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
