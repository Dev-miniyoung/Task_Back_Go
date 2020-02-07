package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model
type Car struct {
	No     string `json:"no"`
	Vin    string `json:"vin"`
	Model  string `json:"model"`
	Make   string `json:"make"`
	Year   int    `json:"year"`
	MSRP   int    `json:"msrp"`
	Status string `json:"status"`
	Booked string `json:"booked"`
	Listed string `json:"listed"`
}

var carData []Car

// Get Inventory data
func getInventoryView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carData)
}

// Get Single Inventory data
func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range carData {
		if item.No == params["no"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

func getCommissionView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Commission Page</h1>")
}

func getManageMarketView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Manage Market Page</h1>")
}

func getManageCustomerView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Manage Customer Page</h1>")
}

func getReportSettingView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Report Setting Page</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, but we couldn't find the page</h1>")
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Dummy Data
	carData = append(carData, Car{
		No:     "1",
		Vin:    "MNBUMF050FW496402",
		Model:  "320i",
		Make:   "BMW",
		Year:   2014,
		MSRP:   147000,
		Status: "ordered",
		Booked: "y",
		Listed: "n",
	})
	carData = append(carData, Car{
		No:     "2",
		Vin:    "4JDBLMF080FW468775",
		Model:  "Camry",
		Make:   "Toyota",
		Year:   2015,
		MSRP:   120000,
		Status: "in stock",
		Booked: "y",
		Listed: "n",
	})
	carData = append(carData, Car{
		No:     "3",
		Vin:    "MNBUMF050FW496402",
		Model:  "320i",
		Make:   "BMW",
		Year:   2014,
		MSRP:   147000,
		Status: "ordered",
		Booked: "y",
		Listed: "n",
	})
	carData = append(carData, Car{
		No:     "4",
		Vin:    "G3SBUMF080FW470449",
		Model:  "Civic",
		Make:   "Honda",
		Year:   2016,
		MSRP:   140000,
		Status: "sold",
		Booked: "n",
		Listed: "n",
	})

	// Route Handlers / Endpoints
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/views/inventory", getInventoryView).Methods("GET")
	r.HandleFunc("/views/inventory/{no}", getCar).Methods("GET")
	r.HandleFunc("/views/commission", getCommissionView).Methods("GET")
	r.HandleFunc("/views/managemarket", getManageMarketView).Methods("GET")
	r.HandleFunc("/views/managecustomer", getManageCustomerView).Methods("GET")
	r.HandleFunc("/views/reportsetting", getReportSettingView).Methods("GET")

	http.ListenAndServe(":8080", r)
}
