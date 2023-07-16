package domain

import (
	"github.com/sjdpk/bankingapp/dto"
	"github.com/sjdpk/bankingapp/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningData string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
