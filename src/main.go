package main

import (
	
	"net/http"
	"encoding/json"
	// "log"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Model
type Car struct{
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

var carData []Car

// Get Inventory data
func getInventory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carData)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Dummy Data
	carData = append(carData, Car{
		No: 1,
		Vin: "MNBUMF050FW496402",
		Model: "320i",
		Make: "BMW",
		Year: 2014,
		MSRP: 147000,
		Status: "ordered",
		Booked: "y",
		Listed: "n",
	})
	carData = append(carData, Car{
		No: 2,
		Vin: "4JDBLMF080FW468775",
		Model: "Camry",
		Make: "Toyota",
		Year: 2015,
		MSRP: 120000,
		Status: "in stock",
		Booked: "y",
		Listed: "n",
	})
	carData = append(carData, Car{
		No: 3,
		Vin: "MNBUMF050FW496402",
		Model: "320i",
		Make: "BMW",
		Year: 2014,
		MSRP: 147000,
		Status: "ordered",
		Booked: "y",
		Listed: "n",
	})
	carData = append(carData, Car{
		No: 4,
		Vin: "G3SBUMF080FW470449",
		Model: "Civic",
		Make: "Honda",
		Year: 2016,
		MSRP: 140000,
		Status: "sold",
		Booked: "n",
		Listed: "n",
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/inventory", getInventory).Methods("GET")
	// r.HandleFunc("/commission", getCommission).Methods("GET")
	// r.HandleFunc("/managemarket", getManageMarket).Methods("GET")
	// r.HandleFunc("/managecustomer", getManageCustomer).Methods("GET")
	// r.HandleFunc("/reportsetting", getReportSetting).Methods("GET")
	
	http.ListenAndServe(":8080", r)
}