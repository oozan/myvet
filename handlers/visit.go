package handlers

import (
	"encoding/json"
	"log"
	"myvet-v2-api/context"
	"myvet-v2-api/structs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//SearchAppointments returns all the appointments for searchterm
func GetWorkList(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/visits/worklist/{date}")
		params := mux.Vars(r)
		dateStr := params["date"]
		_, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Parameter date invalid: " + dateStr})
			return
		}
		res, err := b.Repo.GetWorkList(dateStr)
		//log.Printf("Handlers SearchAppointments res %v", res)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if errRespond(err, w) {
			return
		}
		//log.Printf("Handlers SearchAppointments fileldApts %v", filledAppts)
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

//SearchAppointments returns all the appointments for searchterm
func GetAllReservations(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/visits/reservation/")

		res, err := b.Repo.GetAllReservations()
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

//SearchAppointments returns all the appointments for searchterm
func GetAllReceivables(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/visits/receivable/")

		res, err := b.Repo.GetReceivables()
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}
