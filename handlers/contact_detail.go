package handlers

import (
	"encoding/json"
	"log"
	"myvet-v2-api/context"
	"myvet-v2-api/structs"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateContact handles the /api/create-contact endpoint and adds a row to klinikkaohjelma_kehitys.tyhteystieto
func CreateContact(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-contact/")
		decoder := json.NewDecoder(r.Body)
		var contact structs.ContactDetail
		err := decoder.Decode(&contact)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		err = b.Repo.CreateContact(contact)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
	}
	return fn
}

//GetContactDetails returns
func GetContactDetails(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/contact-details/{customerids}")
		var params = mux.Vars(r)

		idsStr := params["customerids"]
		idStrArr := strings.Split(idsStr, ",")
		var idInts []int
		for _, idStr := range idStrArr {
			i, err := strconv.Atoi(idStr)
			if errRespond(err, w) {
				return
			}
			idInts = append(idInts, i)
		}
		log.Println("Getting contact details for", len(idInts), "customers from the database.")
		retMap, err := b.Repo.GetContactDetails(idInts)
		if errRespond(err, w) {
			return
		}
		log.Println("Returning", len(retMap), "contactdetails.")
		json.NewEncoder(w).Encode(retMap)
	}
	return fn
}

// UpdateContact handles the /api/update-contact endpoint and adds a row to klinikkaohjelma_kehitys.tyhteystieto
func UpdateContact(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/update-contact/")
		decoder := json.NewDecoder(r.Body)
		var contact structs.ContactDetail
		err := decoder.Decode(&contact)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdateContact(contact)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
		return
	}
	return fn
}
