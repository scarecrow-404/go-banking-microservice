package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/scarecrow-404/banking/errs"
)
const (
    host     = "localhost"
    port     = 5432
    user     = "Secret"
    password = "My_Own_Password"
    dbname   = "banking"
)
type CustomerRepositoryDb struct {
	db *sql.DB
}
func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	
	customers := make([]Customer,0)
    
	FindAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err :=  d.db.Query(FindAllSql)
	CheckError(err)
	for rows.Next(){
		var c Customer
		err :=rows.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.DateofBirth,&c.Status)
		if err != nil {
			log.Println("Error while scaning customers:",err.Error())
			return nil ,err
		}
		customers = append(customers,c)
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

	row := d.db.QueryRow(CustomerSql,id)
	var c Customer
	err :=row.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.DateofBirth,&c.Status)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("customer not found")
		}else{
			log.Println("Error while scaning customer:",err.Error())
		return nil,errs.NewUnexpectedError("unexpected database error")
		}
		
	}

	return &c,nil
}



func NewCustomerRepositoryDb() CustomerRepositoryDb{
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
 
    db, err := sql.Open("postgres", psqlconn)
	db.SetConnMaxLifetime(time.Minute *3)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
    CheckError(err)
	
	return CustomerRepositoryDb{db}
}

