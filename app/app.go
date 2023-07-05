package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjdpk/bankingapp/domain"
	"github.com/sjdpk/bankingapp/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
