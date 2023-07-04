package app

import (
	"encoding/json"
	"net/http"

	"github.com/sjdpk/bankingapp/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()
	w.Header().Add("Content-Tye", "application/json")
	json.NewEncoder(w).Encode(customers)
}
