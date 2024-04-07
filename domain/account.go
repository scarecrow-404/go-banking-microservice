package domain

import (
	"github.com/scarecrow-404/banking/dto"
	"github.com/scarecrow-404/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  int  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      int  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	ById(string) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse{
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool{
	return a.Amount > amount
}