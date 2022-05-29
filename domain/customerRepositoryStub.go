package domain

// business - backend adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jimbo", City: "New York", Zipcode: "12213", DateofBirth: "1999-10-19", Status: "1"},
		{Id: "1001", Name: "Rob", City: "San Francisco", Zipcode: "12213", DateofBirth: "2000-02-12", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
