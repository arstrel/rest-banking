package domain

import (
	"errors"
	"fmt"
)

// business - backend adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, error) {
	for _, c := range s.customers {
		if c.Id == id {
			return &c, nil
		}
	}
	msg := fmt.Sprintf("Customer with id: %v is not found", id)
	return nil, errors.New(msg)

}

// Mock Adapter
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jimbo", City: "New York", Zipcode: "12213", DateofBirth: "1999-10-19", Status: "1"},
		{Id: "1002", Name: "Rob", City: "San Francisco", Zipcode: "12213", DateofBirth: "2000-02-12", Status: "1"},
		{Id: "1003", Name: "Stan", City: "Miami", Zipcode: "12213", DateofBirth: "2001-02-12", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
