package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arstrel/rest-banking/domain"
	"github.com/arstrel/rest-banking/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassws := os.Getenv("DB_PASSWD")
	dbPort := os.Getenv("DB_PORT")
	dbAddr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	if address == "" ||
		port == "" ||
		dbUser == "" ||
		dbPassws == "" ||
		dbPort == "" ||
		dbAddr == "" ||
		dbName == "" {
		log.Fatal(fmt.Sprint("Environment variable(s) not defined or missing. Make sure to define:\n",
			"SERVER_ADDRESS\n",
			"SERVER_PORT\n",
			"DB_USER\n",
			"DB_PASSWD\n",
			"DB_PORT\n",
			"DB_ADDR\n",
			"DB_NAME\n"))
	}
}

func Start() {

	sanityCheck()

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

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	// starting server. Second arg is a handler, if nil is passed - default handler will be used
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
