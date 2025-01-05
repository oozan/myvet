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

// CreateTapiolaReply
func CreateTapiolaReply(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-payment/")
		decoder := json.NewDecoder(r.Body)
		var tapiolaReply structs.TapiolaReply
		err := decoder.Decode(&tapiolaReply)
		if errRespond(err, w) {
			return
		}

		res, err := b.Repo.CreateTapiolaReply(tapiolaReply)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

// CreateTapiolaComment
func CreateTapiolaCommmunication(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/tapiola-comment (method POST)")
		decoder := json.NewDecoder(r.Body)
		var tapiolaCommunication structs.TapiolaCommunicationDTO
		err := decoder.Decode(&tapiolaCommunication)
		if err != nil {
			errRespond(err, w)
			return
		}
		if err := b.Repo.CreateTapiolaCommunication(tapiolaCommunication); err != nil {
			errRespond(err, w)
			return
		}
		// empty reply on success?
		json.NewEncoder(w).Encode("{}")
		return
	}
	return fn
}

// CreatePayment handles the /api/create-payment endpoint and adds a row to klinikkaohjelma_kehitys.tsuoritus
func CreatePayment2s(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-payments/")
		decoder := json.NewDecoder(r.Body)
		var pmts []structs.Payment2
		err := decoder.Decode(&pmts)
		if errRespond(err, w) {
			return
		}
		ret, err := b.Repo.CreatePayment2s(pmts)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(ret)
	}
	return fn
}

// for the billnumbers gives as path parameter.
func GetPayment2sByBillnumbers(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/payments2/billnumbers/{billnumbers}")
		params := mux.Vars(r)
		billnumbers := strings.Split(params["billnumbers"], ",")
		var ret []structs.Payment2
		for _, s := range billnumbers {
			i, err := strconv.Atoi(s)
			if errRespond(err, w) {
				return
			}
			recs, err := b.Repo.GetPayment2sByBillNumber(i)
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
