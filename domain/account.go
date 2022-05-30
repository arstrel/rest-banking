package domain

import "github.com/arstrel/rest-banking/errs"

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
