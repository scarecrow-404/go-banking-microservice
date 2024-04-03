package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/scarecrow-404/banking/domain"
	"github.com/scarecrow-404/banking/service"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	// ch:=CustomerHandlers{service:  service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch:=CustomerHandlers{service:  service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define Routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}