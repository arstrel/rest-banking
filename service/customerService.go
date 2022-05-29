package service

import "github.com/arstrel/rest-banking/domain"

// Interactors are classes that orchestrate and perform domain actions —
// think of Service Objects or Use Case Objects.
// They implement complex business rules and validation logic
// specific to a domain action (e.g., onboarding a production)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
