package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer{
		{Id: "001", Name: "Rob", City: "Bangkok", Zipcode: "20000", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "002", Name: "Dabank", City: "CNX", Zipcode: "20000", DateofBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customer}
}