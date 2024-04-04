package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scarecrow-404/banking/service"
)

type CustomerHandlers struct{
	service service.CustomerService
}

func (ch *CustomerHandlers)getAllCustomer(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers,err := ch.service.GetAllCustomer(status)
	if err != nil{
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,customers)
	}
}
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer,err:=ch.service.GetCustomer(id)
	if  err != nil {
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,customer)
		
	}
}

// func (ch *CustomerHandlers) getCustomerByStatus(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	status := vars["status"]
// 	customers,err:=ch.service.GetCustomerByStatus(status)
// 	if  err != nil {
// 		writeResponse(w,err.Code,err.AsMessage())
// 	}else{
// 		writeResponse(w,http.StatusOK,customers)
		
// 	}
// }

// func createCustomer(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w ,"Post request received")
// }

func writeResponse (w http.ResponseWriter, code int,data interface{} ){
	w.Header().Add("Content-Type", "appication/json")
		w.WriteHeader(code)
		if err:= json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
		
}