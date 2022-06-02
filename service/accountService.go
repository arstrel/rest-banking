package service

import (
	"time"

	"github.com/arstrel/rest-banking/domain"
	"github.com/arstrel/rest-banking/dto"
	"github.com/arstrel/rest-banking/errs"
)

// port from domain to the outside world
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

// This holds the reference to a secondary port
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// Here, in the service layer, we do the transformation from DTO to Domain and back as needed
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	acc, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	res := acc.ToNewAccountResponseDto()

	return &res, nil
}

func NewAccountService(r domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{r}
}
