package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"post_code"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}