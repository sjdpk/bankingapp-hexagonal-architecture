package domain

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sjdpk/bankingapp/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	findAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllQuery)
	if err != nil {
		return nil, errs.HandleError(http.StatusInternalServerError, "unexpected database error")
	}
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			return nil, errs.HandleError(http.StatusInternalServerError, "unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findByIdQuery := `select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=$1`
	var c Customer
	row := d.client.QueryRow(findByIdQuery, id)
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.HandleError(http.StatusNotFound, "customer not found")
		} else {
			return nil, errs.HandleError(http.StatusInternalServerError, "unexpected database error")
		}

	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := `postgres://iamdpk:iamdpk@localhost:5432/bankingapp?sslmode=disable`
	client, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{client}
}
