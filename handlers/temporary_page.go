package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"myvet-v2-api/context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// GetTemporaryPage end-point returns either the temporary page by the id, given
// the key provided as query parameter matches the stored one.
// Note: the only type atm is "tapiola", but the t query parameter also must match
// the stored type. A lot of this functionality has to be revamped in the future
// but for now we are cloning the functionality of V1.
func GetTemporaryPage(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/temporary-page/{id}")
		params := mux.Vars(r)
		key := r.URL.Query().Get("k")
		t := r.URL.Query().Get("t")

		if strings.ToLower(t) != "tapiola-reply" {
			err := errors.New("Invalid type parameter")
			errRespond(err, w)
			return
		}

		log.Println(r.URL.Query().Get("k"))
		log.Println(r.URL.Query().Get("t"))
		temporaryPageID, err := strconv.Atoi(params["id"])
		log.Println("Temp page id:", temporaryPageID, "Key:", key, "Type:", t)
		if errRespond(err, w) {
			log.Println("error:" + err.Error())
			return
		}
		// TODO: also use key and type in the query!
		res, err := b.Repo.GetTemporaryPage(temporaryPageID, key, t)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}
