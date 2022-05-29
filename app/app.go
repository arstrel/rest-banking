package app

import (
	"log"
	"net/http"

	"github.com/arstrel/rest-banking/domain"
	"github.com/arstrel/rest-banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	// define custom multiplexer. This is optional because we can use default mux
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// Wiring
	customerRepoMock := domain.NewCustomerRepositoryStub()
	customerRepoDB := domain.NewCustomerRepositoryDb()

	customerServiceMock := service.NewCustomerService(customerRepoMock)
	customerServiceDB := service.NewCustomerService(customerRepoDB)

	customerHandlers := CustomerHandlers{customerServiceDB}
	customerMockHandlers := CustomerHandlers{customerServiceMock}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/mock/customers", customerMockHandlers.getAllCustomers).Methods(http.MethodGet)

	// starting server. Second arg is a handler, if nil is passed - default handler will be used
	log.Fatal(http.ListenAndServe(":8000", router))
}
