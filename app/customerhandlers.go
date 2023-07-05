package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sjdpk/bankingapp/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		WriteResponse(w, err.Code, err)
	} else {
		WriteResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		WriteResponse(w, err.Code, err)
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
