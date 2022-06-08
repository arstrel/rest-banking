package dto

import "github.com/arstrel/rest-banking/errs"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

const WITHDRAWAL string = "Withdrawal"
const DEPOSIT string = "Deposit"

func (t TransactionRequest) Validate() *errs.AppError {
	if t.TransactionType != WITHDRAWAL && t.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}

	if t.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}

func (t TransactionRequest) IsDeposit() bool {
	return t.TransactionType == "Deposit"

}

func (t TransactionRequest) IsWithdrawal() bool {
	return t.TransactionType == "Widthdrawal"
}
