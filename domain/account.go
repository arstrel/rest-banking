package domain

import (
	"github.com/arstrel/rest-banking/dto"
	"github.com/arstrel/rest-banking/errs"
)

// Account domain object
type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// Secondary port - interface
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
