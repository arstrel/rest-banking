package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/arstrel/rest-banking/errs"
	"github.com/arstrel/rest-banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error while querying customer table")
	}

	// sqlx Select combines querying and marshaling so there is no need for this line anymore
	// err = sqlx.StructScan(rows, &customers)

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	c := Customer{}
	err := d.client.Get(&c, findByIdSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &c, nil
}

// Create a connection
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPassws := os.Getenv("DB_PASSWD")
	dbPort := os.Getenv("DB_PORT")
	dbAddr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassws, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
