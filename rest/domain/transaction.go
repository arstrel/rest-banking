package domain

import (
	"github.com/arstrel/rest-banking/rest/dto"
	"github.com/arstrel/rest-banking/rest/errs"
)

type Transaction struct {
	TransactionId string  `json:"transaction_id"`
	AccountId     string  `json:"account_id"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"transaction_type"`
	Date          string  `json:"transaction_date"`
}

func (t Transaction) ToDto() (*dto.TransactionResponse, *errs.AppError) {
	res := dto.TransactionResponse{
		TransactionId: t.TransactionId,
		AccountId:     t.AccountId,
		Amount:        t.Amount,
		Type:          t.Type,
		Date:          t.Date,
	}

	return &res, nil
}

func (t Transaction) IsWithdrawal() bool {
	return t.Type == "Widthdrawal"
}
