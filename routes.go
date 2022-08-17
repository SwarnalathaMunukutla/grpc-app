package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1" xml:"addressline1"`
	AddressLine2 string `json:"address_line_2" xml:"addressline2"`
	City         string `json:"city" xml:"city"`
	State        string `json:"state" xml:"state"`
	Country      string `json:"country" xml:"country"`
	Zipcode      int64  `json:"zip_code" xml:"zipcode"`
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashish", City: "New Delhi", Zipcode: "110075"},
		{Name: "Anish", City: "New Delhi", Zipcode: "110025"},
		{Name: "Madhav", City: "Chennai", Zipcode: "410075"},
		{Name: "Rajesh", City: "Patna", Zipcode: "210075"},
		{Name: "Jaykumar", City: "New Delhi", Zipcode: "530002"},
		{Name: "Veena", City: "Mumbai", Zipcode: "200015"},
		{Name: "Anita", City: "Pune", Zipcode: "450015"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func getFullAddress(w http.ResponseWriter, r *http.Request) {
	address := Address{AddressLine1: " A/902, Beverly Hills", AddressLine2: "Baner", City: "Pune", State: "Maharashtra", Country: "India", Zipcode: 530002}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(address)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(address)
	}
}
