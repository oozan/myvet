package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"myvet-v2-api/context"
	"myvet-v2-api/structs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetWeightHistory returns the weight history for the given animal
func GetWeightHistory(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("handlers.GetWeightHistory")
		var res []structs.WeightHistoryEntry
		params := mux.Vars(r)
		animalID := params["animalID"]
		animalIDInt, err := strconv.Atoi(animalID)
		if err != nil {
			log.Println("GetWeightHistory error: " + err.Error())
			err = errors.New("Invalid animal id")
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		if res, err = b.Repo.GetWeightHistoryForAnimal(animalIDInt); err != nil {
			log.Println("GetWeightHistory error: " + err.Error())
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}
