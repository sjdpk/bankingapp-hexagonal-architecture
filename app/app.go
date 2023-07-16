package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sjdpk/bankingapp/domain"
	"github.com/sjdpk/bankingapp/service"
)

func sanityCheck() {
	// server addr check
	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")
	// db env variable
	dbDriver := os.Getenv("DB")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if serverAddr == "" || serverPort == "" {
		log.Fatal("Server Environment variable not defined...")
	} else if dbDriver == "" || dbUser == "" || dbPasswd == "" || dbAddr == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database Environment variable not defined...")
	}
}
func Start() {
	//environment variable check
	sanityCheck()
	router := mux.NewRouter()
	// wiring
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandler{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customer", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id}", ch.getCustomer).Methods(http.MethodGet)
	// account handler
	router.HandleFunc("/customer/{customer_id}/account", ah.NewAccount).Methods(http.MethodPost)

	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")
	serverAddrPort := fmt.Sprintf("%s:%s", serverAddr, serverPort)

	log.Fatal(http.ListenAndServe(serverAddrPort, router))
}

func getDbClient() *sqlx.DB {
	dbDriver := os.Getenv("DB")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbDriver, dbUser, dbPasswd, dbAddr, dbPort, dbName)
	// connStr := `postgres://iamdpk:iamdpk@localhost:5432/bankingapp?sslmode=disable`
	connStr := dbUrl + `?sslmode=disable`
	client, err := sqlx.Open(dbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
