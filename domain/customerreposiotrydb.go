package domain

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sjdpk/bankingapp/errs"
	"github.com/sjdpk/bankingapp/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var (
		// rows *sql.Rows
		err error
	)
	customers := make([]Customer, 0)

	if status != "" {
		findAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status= $1"
		err = d.client.Select(&customers, findAllQuery, status)
	} else {
		findAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllQuery)

	}
	if err != nil {
		logger.Error("unexpected database error: " + err.Error())
		return nil, errs.HandleError(http.StatusInternalServerError, "unexpected database error")
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findByIdQuery := `select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=$1`
	var c Customer
	err := d.client.Get(&c, findByIdQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.HandleError(http.StatusNotFound, "customer not found")
		} else {
			logger.Error("unexpected database error: " + err.Error())
			return nil, errs.HandleError(http.StatusInternalServerError, "unexpected database error")
		}

	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
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
	return CustomerRepositoryDB{client}
}
