package domain

import (
	"github.com/arstrel/rest-banking/dto"
	"github.com/arstrel/rest-banking/errs"
)

// Account domain object
type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

// Secondary port - interface
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindByAccountId(id string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
