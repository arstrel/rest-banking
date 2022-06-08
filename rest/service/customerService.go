package service

import (
	"github.com/arstrel/rest-banking/rest/domain"
	"github.com/arstrel/rest-banking/rest/dto"
	"github.com/arstrel/rest-banking/rest/errs"
)

// Interactors are classes that orchestrate and perform domain actions —
// think of Service Objects or Use Case Objects.
// They implement complex business rules and validation logic
// specific to a domain action (e.g., onboarding a production)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	statusToCodeMap := map[string]string{
		"active":   "1",
		"inactive": "0",
	}

	if sCode, ok := statusToCodeMap[status]; ok {
		status = sCode
	} else {
		status = ""
	}

	cust, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	var resp = make([]dto.CustomerResponse, 0)

	for _, c := range cust {
		resp = append(resp, c.ToDto())
	}

	return resp, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
