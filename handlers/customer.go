package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"myvet-v2-api/context"
	"myvet-v2-api/structs"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateCustomer handles the /api/create-customer endpoint and adds a row to klinikkaohjelma_kehitys.tasiakas
func CreateCustomer(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-customer/")
		decoder := json.NewDecoder(r.Body)
		var customer structs.Customer
		err := decoder.Decode(&customer)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		_, err = b.Repo.CreateCustomer(customer)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
	}
	return fn
}

// api/search/customers/{terms}
func SearchCustomers(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/search/customers/{terms}")
		params := mux.Vars(r)
		termStr := params["terms"]
		termsArr := strings.Split(termStr, ",")
		if len(termsArr) == 0 {
			err := errors.New("No search terms given")
			log.Printf("SearchCustomers error: %s\n", err.Error())
			errRespond(err, w)
			return
		}
		for i, t := range termsArr {
			termsArr[i] = strings.ToLower(t)
		}
		res, err := b.Repo.SearchCustomers(termsArr)
		if err != nil {
			errRespond(err, w)
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

//SearchCustomersAnimals commits searches against the customers and animals tables (telain, tasiakas).
func SearchCustomersAnimals(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/search/customersAnimals/{terms}")
		params := mux.Vars(r)
		termStr := params["terms"]
		termArr := strings.Split(termStr, ",")
		res, err := b.Repo.SearchCustomersAnimals(termArr)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

//SearchCustomersAnimals commits searches against the customers and animals tables (telain, tasiakas).
func GetCustomerAnimals(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/customer/{customerId}/animals")
		params := mux.Vars(r)
		idStr := params["customerId"]
		log.Println("PPÄÄÄÄ" + idStr)
		customerID, err := strconv.Atoi(idStr)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		res, err := b.Repo.GetCustomerAnimals(customerID)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

// GetAllCustomers responds with a slice of customers, limited only by offset/limit.
func GetAllCustomers(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/customers/")

		offset := 100
		limit := 150
		//offset := 100
		//limit := 150
		//log.Fatal("Handling /api/customers/  jeees")
		res, err := b.Repo.GetAllCustomers(offset, limit)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

//GetCustomer handles one customer information by ID.
func GetCustomer(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/customer/{id}")
		params := mux.Vars(r)
		intID, err := strconv.Atoi(params["id"])
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetCustomer(intID)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

//GetCustomersByIDs handles /customers/id/
func GetCustomersByIDs(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling GET /customers/ids/{ids}")
		params := mux.Vars(r)
		idsStr := params["ids"]
		idStrArr := strings.Split(idsStr, ",")
		var idInts []int
		for _, idStr := range idStrArr {
			i, err := strconv.Atoi(idStr)
			if errRespond(err, w) {
				return
			}
			idInts = append(idInts, i)
		}
		log.Println("Getting", len(idInts), "customers from the database.")

		retCustomers, err := b.Repo.GetCustomers(idInts)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(retCustomers)
	}

	return fn
}

// UpdateCustomer handles the /api/update-customer endpoint and adds a row to klinikkaohjelma_kehitys.tasiakas
func UpdateCustomer(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/update-customer/")
		decoder := json.NewDecoder(r.Body)
		var customer structs.Customer
		err := decoder.Decode(&customer)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdateCustomer(customer)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
	}
	return fn
}
