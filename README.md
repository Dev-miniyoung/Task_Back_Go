## CodeBrick_Back_Assignment

### Task Summary

1. use Go
2. create single page web application with ajax
3. use router (http://www.gorillatoolkit.org/)

### Completed Function

1. api endpoint
2. use router

### API DOCS

**Method:**
`GET`

**Sample Call:**

```
func getInventoryView(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(carData)
}


r.HandleFunc("/inventory", getInventory).Methods("GET")
Success Response:
Code: 200
Content: {
No:     "1",
Vin:    "MNBUMF050FW496402",
Model:  "320i",
Make:   "BMW",
Year:   2014,
MSRP:   147000,
Status: "ordered",
Booked: "y",
Listed: "n",
},{
No:     "2",
Vin:    "4JDBLMF080FW468775",
Model:  "Camry",
Make:   "Toyota",
Year:   2015,
MSRP:   120000,
Status: "in stock",
Booked: "y",
Listed: "n",
}...,


r.HandleFunc("/inventory/{no}", getCar).Methods("GET")
ex) http://localhost:8080/views/inventory/1
Success Response:
Code: 200
Content: {
No:     "1",
Vin:    "MNBUMF050FW496402",
Model:  "320i",
Make:   "BMW",
Year:   2014,
MSRP:   147000,
Status: "ordered",
Booked: "y",
Listed: "n",
}

r.HandleFunc("/views/commission", getCommissionView).Methods("GET")
r.HandleFunc("/views/managemarket", getManageMarketView).Methods("GET")
r.HandleFunc("/views/managecustomer", getManageCustomerView).Methods("GET")
r.HandleFunc("/views/reportsetting", getReportSettingView).Methods("GET")
```
