package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arstrel/rest-banking/domain"
	"github.com/arstrel/rest-banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbPort := os.Getenv("DB_PORT")
	dbAddr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	// "#{dbUser}:#{dbPasswd}@tcp(#{dbAddr}:#{dbPort})/#{dbName}"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

func Start() {

	sanityCheck()

	// define custom multiplexer. This is optional because we can use default mux
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	dbClient := getDbClient()

	// Wiring
	customerRepoDB := domain.NewCustomerRepositoryDb(dbClient)
	customerServiceDB := service.NewCustomerService(customerRepoDB)
	customerHandlers := CustomerHandlers{customerServiceDB}

	accountRepoDB := domain.NewAccountRepositoryDb(dbClient)
	accountServiceDB := service.NewAccountService(accountRepoDB)
	accountHandlers := AccountHandlers{accountServiceDB}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandlers.getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}/account", accountHandlers.newAccount).Methods(http.MethodPost)

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
