package domain

import "github.com/arstrel/rest-banking/errs"

// Entities are the domain objects (Business object)
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

// Repositories are the interfaces to getting entities as well as creating and changing them.
// They keep a list of methods that are used to communicate with data sources
// and return a single entity or a list of entities. (e.g. UserRepository)

// Repository interface - secondary port. Sits in between Domain(Business) and Backend(ServerSide)
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
