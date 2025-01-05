package handlers

import (
	"encoding/json"
	"myvet-v2-api/context"
	"net/http"
)

/*
 * GETVatRates is a handler function for /api/vat-rates
 * and returns the whole history of used VAT rates and categories in finland
 */
func GetVATRates(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		res, err := b.Repo.GetVATHistory()
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}
