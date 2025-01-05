package handlers

import (
	"encoding/json"
	"log"
	"myvet-v2-api/context"
	"net/http"

	"github.com/gorilla/mux"
)

//GetReceivablesByDate returns list of work shift information given by date.
func GetReceivablesByDate(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/receivables/date/{date}")
		params := mux.Vars(r)
		strDate := params["date"]
		if len(strDate) != 10 {
			strErrRespond("date required in YYYY-MM-DD format.", w)
			return
		}
		res, err := b.Repo.GetReceivablesByDate(strDate)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}
