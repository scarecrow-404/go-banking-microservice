package domain

import (
	"github.com/scarecrow-404/banking/dto"
	"github.com/scarecrow-404/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string 
	City        string
	Zipcode     string
	DateofBirth string  `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAstext() string {
	statusAstext := "active"
	if c.Status == "0" {
		statusAstext = "inactive"
	}
	return statusAstext
}

func (c Customer) ToDto() dto.CustomerResponse {
	
	return dto.CustomerResponse{
		Id : c.Id,
		Name: c.Name,
		City: c.City,
		Zipcode: c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status: c.statusAstext(),
	}

}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)

	// FindStatus(string) ([]Customer, *errs.AppError)
}
