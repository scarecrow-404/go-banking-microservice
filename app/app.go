package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/scarecrow-404/banking/domain"
	"github.com/scarecrow-404/banking/logger"
	"github.com/scarecrow-404/banking/service"
)

func sanityCheck() {
	envVariable :=[]string{
		"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASSWD", "DB_HOST", "DB_PORT", "DB_NAME",
	}
	for _, v := range envVariable {
		if os.Getenv(v) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", v))
		}
	}
}

func Start() {

	sanityCheck()
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	// ch:=CustomerHandlers{service:  service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch:=CustomerHandlers{service:  service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define Routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers?status={status}", ch.getCustomerByStatus).Methods(http.MethodGet)
	//starting server
	address:=os.Getenv("SERVER_ADDRESS")
	port:= os.Getenv("SERVER_PORT")
	addressPort := fmt.Sprintf("%s:%s", address, port)
	fmt.Println("Starting server on", addressPort)
	log.Fatal(http.ListenAndServe(addressPort, router))
	
}