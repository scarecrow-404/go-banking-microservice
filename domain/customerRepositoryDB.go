package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/scarecrow-404/banking/errs"
	"github.com/scarecrow-404/banking/logger"
)
const (
    host     = "localhost"
    port     = 5432
	user     = "Secret"
    password = "My_Own_Password"
    dbname   = "for testing only"
)
type CustomerRepositoryDb struct {
	db *sqlx.DB
}
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer,0)
	var err error
    if status == ""{
		FindAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err =d.db.Select(&customers,FindAllSql)
		
	}else {
		FindAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = $1"
		err =d.db.Select(&customers,FindAllSql,status)
	}
	if err != nil{
		logger.Error("Error while querying customers:"+err.Error())
			return nil, errs.NewUnexpectedError ("unexpected database error")
	}
	return customers, nil
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func (d  CustomerRepositoryDb) ById(id string)( *Customer,*errs.AppError){
	CustomerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"
	var c Customer
	err := d.db.Get(&c,CustomerSql,id)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("customer not found")
		}else{
			logger.Error("Error while scaning customer:"+err.Error())
		return nil,errs.NewUnexpectedError("unexpected database error")
		}
		
	}
	return &c,nil
}




func NewCustomerRepositoryDb() CustomerRepositoryDb{
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
    db, err := sqlx.Open("postgres", psqlconn)
	db.SetConnMaxLifetime(time.Minute *3)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
    CheckError(err)
	
	
	return CustomerRepositoryDb{db}
}

