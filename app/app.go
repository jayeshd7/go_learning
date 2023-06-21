package app

import (
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
