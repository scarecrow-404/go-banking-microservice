package domain

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
		CheckError(err)
		customers = append(customers,c)
	}
	defer d.db.Close()
	return customers, nil
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func NewCustomerRepositoryDb() CustomerRepositoryDb{
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
 
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
	return CustomerRepositoryDb{db}
}