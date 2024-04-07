package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scarecrow-404/banking/dto"
	"github.com/scarecrow-404/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	
	vars:=mux.Vars(r)
	customerId := vars["customer_id"]
	customerIdInt, err := strconv.Atoi(customerId)
	if err!= nil {
    	writeResponse(w, http.StatusBadRequest, "Invalid customer ID")
    	return
	}
	var request dto.NewAccountRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err!= nil {
  	  writeResponse(w, http.StatusBadRequest, err.Error())
  	  return
	}
	request.CustomerId = customerIdInt
	response, appError := h.service.NewAccount(request)
	if appError!= nil {
  	  writeResponse(w, appError.Code, appError.AsMessage())
	  return
	}
	writeResponse(w, http.StatusCreated, response)
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars :=  mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err!= nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.AccountId = accountId
	if err!= nil {
		writeResponse(w, http.StatusBadRequest, "invalid account id")
		return
	}
	request.CustomerId= customerId
	if err!= nil {
		writeResponse(w, http.StatusBadRequest, "invalid customer id")
		return
	}
	response, appError := h.service.MakeTransaction(request)
	if appError!= nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, response)

}