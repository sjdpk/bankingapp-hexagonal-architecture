package dto

import (
	"net/http"
	"strings"

	"github.com/sjdpk/bankingapp/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.HandleError(http.StatusUnprocessableEntity, "to open new account you have to deposits at-least 5000")
	} else if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.HandleError(http.StatusUnprocessableEntity, "account type should be checking or saving")
	}
	return nil
}
