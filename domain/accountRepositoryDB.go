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

func NewAccountREpositoryDB(dbClient *sqlx.DB) AccountRepositoryDB{
	return AccountRepositoryDB{dbClient}
}

