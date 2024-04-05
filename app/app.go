package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

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
	dbClient:= getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountREpositoryDB(dbClient)
	ch:=CustomerHandlers{service:  service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}
	//define Routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	//starting server
	address:=os.Getenv("SERVER_ADDRESS")
	port:= os.Getenv("SERVER_PORT")
	addressPort := fmt.Sprintf("%s:%s", address, port)
	fmt.Println("Starting server on", addressPort)
	log.Fatal(http.ListenAndServe(addressPort, router))
	
}

func getDbClient() *sqlx.DB{
	host     := os.Getenv("DB_HOST")
    port     := os.Getenv("DB_PORT")
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWD")
	dbname   := os.Getenv("DB_NAME")


	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
    db, err := sqlx.Open("postgres", psqlconn)
	if err !=nil{
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute *3)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	return db
}

