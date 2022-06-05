package domain

import (
	"database/sql"
	"strconv"

	"github.com/arstrel/rest-banking/errs"
	"github.com/arstrel/rest-banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while saving an account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last inserted id for the account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDb) FindByAccountId(id string) (*Account, *errs.AppError) {
	sqlFindById := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"

	a := Account{}
	err := d.client.Get(&a, sqlFindById, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		}
		logger.Error("Error while scanning account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) SaveTransaction(tr Transaction) (*Transaction, *errs.AppError) {
	// Starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// inserting bank account transaction

	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) value (?, ?, ?, ?)", tr.AccountId, tr.Amount, tr.Type, tr.Date)

	// Updating account balance
	if tr.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", tr.Amount, tr.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", tr.Amount, tr.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// Commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while saving the transaction and updating the account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// Getting the latest account information from the accounts table
	acc, appErr := d.FindByAccountId(tr.AccountId)

	if appErr != nil {
		return nil, appErr
	}

	tr.TransactionId = strconv.FormatInt(id, 10)

	// updating the transaction struct with the latest balance
	tr.Amount = acc.Amount

	return &tr, nil
}
