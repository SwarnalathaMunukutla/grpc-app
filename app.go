package main

import (
	"log"
	"net/http"
)

func start() {

	// Hello world, the web server
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/customers", getAllCustomersHandler)
	http.HandleFunc("/address", getFullAddress)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
