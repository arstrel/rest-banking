package app

import (
	"log"
	"net/http"
)

func Start() {
	// define custom multiplexer. This is optional because we can use default mux
	mux := http.NewServeMux()
	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server. Second arg is a handler, if nil is passed - default handler will be used
	log.Fatal(http.ListenAndServe(":8000", mux))
}
