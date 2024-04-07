package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/scarecrow-404/banking/errs"
	"github.com/scarecrow-404/banking/logger"
)

type AccountRepositoryDB struct {
	db *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account)(*Account,*errs.AppError){
	var id int64
	
	sqlInsert := "insert into accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) returning account_id"
	err := d.db.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)
	if err != nil {
		logger.Error("Error while creating new account:"+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}
	a.AccountId=strconv.FormatInt(id,10)
	return &a,nil
}

func (d AccountRepositoryDB) ById(id string)( *Account,*errs.AppError){
	sqlSelect := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = $1"
	var a Account
	err := d.db.Get(&a,sqlSelect,id)
	if err != nil {
			logger.Error("Error while fetching an account:"+err.Error())
			return nil,errs.NewUnexpectedError("unexpected database error")
		}
	return &a,nil
}

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction,*errs.AppError){
	tn,err := d.db.Begin()
	if err != nil {
		logger.Error("Error while starting new transaction : "+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}
	err = tn.QueryRow("insert into transactions (account_id, amount, transaction_type, transaction_date) values ($1, $2, $3, $4) returning transaction_id",t.AccountId,t.Amount,t.TransactionType,t.TransactionDate).Scan(&t.TransactionId)
	if err != nil {
		logger.Error("Error while creating new transaction : "+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}
	if t.IsWithdraw() {
		_,err = tn.Exec("update accounts set amount = amount - $1 where account_id = $2",t.Amount,t.AccountId)
	} else {
		_,err = tn.Exec("update accounts set amount = amount + $1 where account_id = $2",t.Amount,t.AccountId)
	}
	if err != nil{
		tn.Rollback()
		logger.Error("Error while saving transaction : "+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}

	err = tn.Commit()
	if err != nil {
		tn.Rollback()
		logger.Error("Error while commiting transaction : "+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}
	acc,appErr := d.ById(t.AccountId)
	if appErr != nil {
		return nil,appErr
	}

	t.Amount = acc.Amount
	return &t,nil
}

func NewAccountREpositoryDB(dbClient *sqlx.DB) AccountRepositoryDB{
	return AccountRepositoryDB{dbClient}
}

