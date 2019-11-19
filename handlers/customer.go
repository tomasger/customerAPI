package handlers

import (
	"customerAPI/models"
	"customerAPI/storage"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var PostCustomer = func(db storage.Database, w http.ResponseWriter, r *http.Request) {
	var customer *models.Customer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		w.WriteHeader(400)
		return
	}
	defer r.Body.Close()

	// check data validity
	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		w.WriteHeader(400)
		return
	}


	customer.Date = time.Now().String()

	if err := db.CreateCustomer(customer); err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}

var GetCustomer = func(db storage.Database, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// mux ensures to only accept integers as ID
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	customer, err := db.GetCustomer(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

var GetCustomers = func(db storage.Database, w http.ResponseWriter, r *http.Request) {
	customers, err := db.GetCustomers()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
