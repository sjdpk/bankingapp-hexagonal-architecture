package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjdpk/bankingapp/domain"
	"github.com/sjdpk/bankingapp/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello THere")
}
