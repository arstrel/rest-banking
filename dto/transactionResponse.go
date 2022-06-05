package dto

type TransactionResponse struct {
	TransactionId string
	AccountId     string
	Amount        float64
	Type          string
	Date          string
}
