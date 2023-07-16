package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/sjdpk/bankingapp/errs"
	"github.com/sjdpk/bankingapp/logger"
	"net/http"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) VALUES ($1,$2,$3,$4,$5) RETURNING account_id"
	var accountId string
	err := d.client.QueryRow(sqlInsert, a.CustomerId, a.OpeningData, a.AccountType, a.Amount, a.Status).Scan(&accountId)
	if err != nil {
		logger.Error("Error While Creating new account: " + err.Error())
		return nil, errs.HandleError(http.StatusInternalServerError, "unexpected error from database")
	}
	a.AccountId = accountId
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
