package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sjdpk/bankingapp/domain"
	"github.com/sjdpk/bankingapp/service"
)

func sanityCheck() {
	// server addr check
	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")
	// db env varibale
	dbDriver := os.Getenv("DB")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if serverAddr == "" || serverPort == "" {
		log.Fatal("Server Environment varibale not defined...")
	} else if dbDriver == "" || dbUser == "" || dbPasswd == "" || dbAddr == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database Environment varibale not defined...")
	}
}
func Start() {
	//environment variable check
	sanityCheck()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/customer", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")
	serverAddrPort := fmt.Sprintf("%s:%s", serverAddr, serverPort)

	log.Fatal(http.ListenAndServe(serverAddrPort, router))
}
