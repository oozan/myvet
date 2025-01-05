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

// CreateAnimal handles the /api/create-animal endpoint and adds a row to klinikkaohjelma_kehitys.tlaji
func CreateAnimal(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-animal/")
		decoder := json.NewDecoder(r.Body)
		var animal structs.Animal
		err := decoder.Decode(&animal)
		if err != nil {
			log.Println("CreateAnimal decoding error: " + err.Error())
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		log.Println("Handlers.CreateAnimal animal.Name " + animal.Name.String)
		log.Println("Handlers.CreateAnimal animal.CustomerID " + strconv.Itoa(animal.CustomerID))

		id, err := b.Repo.CreateAnimal(animal)
		if err != nil {
			log.Println("CreateAnimal error: " + err.Error())
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		// TODO: get the animal
		resAnimal, errAnimal := b.Repo.GetAnimal(id)
		if errAnimal != nil {
			log.Println("CreateAnimal error: " + errAnimal.Error())
			errRespond(errAnimal, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resAnimal)
		return
	}
	return fn
}

// UpdateAnimal handles the /api/update-animal endpoint and adds a row to klinikkaohjelma_kehitys.telain
func UpdateAnimal(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		animalID, errConv := strconv.Atoi(params["animalID"])
		if errRespond(errConv, w) {
			return
		}
		log.Println("Handling /api/animals/" + strconv.Itoa(animalID))
		decoder := json.NewDecoder(r.Body)
		var animal structs.Animal
		err := decoder.Decode(&animal)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdateAnimal(animal)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
	}
	return fn
}

//GetAnimal handles one animal information by ID.
func GetAnimal(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/animal/{id}")
		params := mux.Vars(r)
		animalID, err := strconv.Atoi(params["id"])
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetAnimal(animalID)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

//GetAnimalsByIDs han§dles /animals/ids/{ids} where the path parameter is comma separated list of ints.
func GetAnimalsByIDs(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling GET /animals/ids/{ids}")
		var params = mux.Vars(r)

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
		log.Println("Getting", len(idInts), "animals from the database.")
		retAnimals, err := b.Repo.GetAnimals(idInts)
		if errRespond(err, w) {
			return
		}
		log.Println("Returning", len(retAnimals), "animals.")
		json.NewEncoder(w).Encode(retAnimals)
		return
	}

	return fn
}

// GetSpeciesByID responds with a Species for the query param id.
func GetSpeciesByID(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/species/{id}")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetSpeciesByID(uint32(id))
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

//GetBreedsBySpecies returns all species by id
func GetBreedsBySpecies(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/breeds/species/{speciesid}")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetSpeciesByID(uint32(id))
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		if res != nil {
			json.NewEncoder(w).Encode(res)

		} else {
			json.NewEncoder(w).Encode([]int{})
			return
		}
		return
	}
	return fn
}

// GetAllSpecies responds with a slice of all Species in the system.
func GetAllSpecies(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/species/")

		res, err := b.Repo.GetAllSpecies()
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

// GetAllBreeds responds with a slice of all Breeds in the system.
func GetAllBreeds(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/breeds/")

		res, err := b.Repo.GetAllBreeds()
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}
