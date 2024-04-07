package domain

import "github.com/scarecrow-404/banking/dto"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

const WITHDRAW = "withdraw"
const DEPOSIT = "deposit"

func (t Transaction) IsWithdraw() bool {
	return t.TransactionType == WITHDRAW
}

func (t Transaction) IsDeposit() bool {
	return t.TransactionType == DEPOSIT
}

func (t Transaction) ToDto() dto.TransactionResponse{
	return dto.TransactionResponse{
		TransactionId: t.TransactionId,
		AccountId: t.AccountId,
		Amount: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}