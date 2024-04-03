package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scarecrow-404/banking/service"
)

type Customer struct{
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city_name"`
	Zipcode string  `json:"post_code"  xml:"postcode"`
}

type CustomerHandlers struct{
	service service.CustomerService
}

func (ch *CustomerHandlers)getAllCustomer(w http.ResponseWriter, r *http.Request) {

	customers,_ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "appication/json")
		json.NewEncoder(w).Encode(customers)
	}

}
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer,err:=ch.service.GetCustomer(id)
	if  err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w,err.Message)
	}else{
		w.Header().Add("Content-Type", "appication/json")
		json.NewEncoder(w).Encode(customer)
	}
}


func createCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w ,"Post request received")
}