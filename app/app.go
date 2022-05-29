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
	customerRepo := domain.NewCustomerRepositoryStub()
	customerService := service.NewCustomerService(customerRepo)
	customerHandlers := CustomerHandlers{customerService}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)

	// starting server. Second arg is a handler, if nil is passed - default handler will be used
	log.Fatal(http.ListenAndServe(":8000", router))
}
