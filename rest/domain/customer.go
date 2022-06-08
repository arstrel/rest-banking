package domain

import (
	"github.com/arstrel/rest-banking/rest/dto"
	"github.com/arstrel/rest-banking/rest/errs"
)

// This is a server side layer

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
	// status == "1" | "0" | ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

// With this function domain has an ability to convert Customer (domain object) to DTO (data transfer object)
func (c Customer) ToDto() dto.CustomerResponse {

	resp := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		DateofBirth: c.DateofBirth,
		Zipcode:     c.Zipcode,
		Status:      c.statusAsText(),
	}

	return resp
}
