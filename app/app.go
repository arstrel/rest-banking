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
	customerRepoDB := domain.NewCustomerRepositoryDb()
	customerServiceDB := service.NewCustomerService(customerRepoDB)
	customerHandlers := CustomerHandlers{customerServiceDB}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandlers.getCustomerById).Methods(http.MethodGet)

	// Wiring Mock
	customerRepoMock := domain.NewCustomerRepositoryStub()
	customerServiceMock := service.NewCustomerService(customerRepoMock)
	customerMockHandlers := CustomerHandlers{customerServiceMock}
	// define mock routes
	router.HandleFunc("/mock/customers", customerMockHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/mock/customers/{id:[0-9]+}", customerMockHandlers.getCustomerById).Methods(http.MethodGet)

	// starting server. Second arg is a handler, if nil is passed - default handler will be used
	log.Fatal(http.ListenAndServe(":8000", router))
}
