package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Jayesh!!")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"jayesh", "pune", "411057"},
		{"pooja", "pune", "411057"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
