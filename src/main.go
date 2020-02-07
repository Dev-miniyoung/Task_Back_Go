package main

import (
	"fmt" 
	"net/http"

	"github.com/gorilla/mux"
)

// Model
type Inventory struct{
	No 		int 	`json:"no"`
	Vin 	string 	`json:"vin"`
	Model 	string 	`json:"model"`
	Make 	string 	`json:"make"`
	Year 	int 	`json:"year"`
	MSRP 	int 	`json:"msrp"`
	Status 	string 	`json:"status"`
	Booked 	string 	`json:"booked"`
	Listed 	string 	`json:"listed"`
}

func main(){
	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/inventory", getInventory).Methods("GET")
	r.HandleFunc("/commission", getCommission).Methods("GET")
	r.HandleFunc("/managemarket", getManageMarket).Methods("GET")
	r.HandleFunc("/managecustomer", getManageCustomer).Methods("GET")
	r.HandleFunc("/reportsetting", getReportSetting).Methods("GET")
	
	http.ListenAndServe(":8080", r)
}