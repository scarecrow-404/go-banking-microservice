package dto

import (
	"github.com/scarecrow-404/banking/errs"
)

type TransactionRequest struct {
	CustomerId string  `json:"customer_id"`
	AccountId   string  `json:"account_id"`
	TransactionDate  string  `json:"transaction_date"`
	Amount      float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
}

type TransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	AccountId   string  `json:"account_id"`
	TransactionDate  string  `json:"transaction_date"`
	Amount      float64 `json:"new_balance"`
	TransactionType string `json:"transaction_type"`
}

const WITHDRAW = "withdraw"
const DEPOSIT = "deposit"

func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != "withdraw" && r.TransactionType != "deposit" {
		return errs.NewValidationError("Invalid transaction type")
	}

	if r.Amount <= 0 {
		return errs.NewValidationError("Invalid amount")
	}
	return nil
}

func (r TransactionRequest) IsWithdraw() bool {
	return r.TransactionType == "withdraw"
}

func (r TransactionRequest) IsDeposit() bool {
	return r.TransactionType == "deposit"
}

