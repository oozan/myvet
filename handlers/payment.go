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

// CreatePayment handles the /api/create-payment endpoint and adds a row to klinikkaohjelma_kehitys.tsuoritus
func CreatePayment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-payment/")
		decoder := json.NewDecoder(r.Body)
		var payment structs.Payment
		err := decoder.Decode(&payment)
		if errRespond(err, w) {
			return
		}
		lastID, err := b.Repo.CreatePayment(payment)
		if errRespond(err, w) {
			return
		}
		payment.PaymentID = int(lastID)
		json.NewEncoder(w).Encode(payment)
		return
	}
	return fn
}

// UpdatePayment handles the /api/update-payment endpoint and adds a row to klinikkaohjelma_kehitys.tsuoritus
func UpdatePayment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/update-payment/")
		decoder := json.NewDecoder(r.Body)
		var payment structs.Payment
		err := decoder.Decode(&payment)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdatePayment(payment)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
		return
	}
	return fn
}

// GetPaymentsByBillnumbers responds with a slice of Payment
// for the billnumbers gives as path parameter.
func GetPaymentsByBillnumbers(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/payments/billnumbers/{billnumbers}")
		params := mux.Vars(r)
		billnumbers := strings.Split(params["billnumbers"], ",")
		var ret []structs.Payment
		for _, s := range billnumbers {
			i, err := strconv.Atoi(s)
			if errRespond(err, w) {
				return
			}
			recs, err := b.Repo.GetPaymentsByBillNumber(i)
			if errRespond(err, w) {
				return
			}
			ret = append(ret, recs...)
		}
		json.NewEncoder(w).Encode(ret)
		return
	}
	return fn
}

// FindNonMatchingSum responds with a list of payments. (See Repo.FindFaultyPayment for more information.)
func FindNonMatchingSum(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/debug2/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.FindNonMatchingSum(strStart, strEnd)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

// FindErroneousPayments responds with a list of payments. (See Repo.FindErroneousPayments for more information.)
func FindErroneousPayments(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/debug1/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.FindErroneousPayments(strStart, strEnd)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

//GetTotalPayment returns list of total payment information given by date.
func GetTotalPayment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/total-payment/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			strErrRespond("Start and end query parameters required in YYYY-MM-DD format, got "+strStart+" and "+strEnd, w)
			return
		}
		err := b.Repo.GetTotalPayment(strStart, strEnd)
		if errRespond(err, w) {
			return
		}
	}
	return fn
}
