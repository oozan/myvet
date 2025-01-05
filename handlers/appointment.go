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

func CreateAppointment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-appointment/")
		decoder := json.NewDecoder(r.Body)
		var appointment structs.Appointment
		err := decoder.Decode(&appointment)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		_, err = b.Repo.CreateAppointment(appointment)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
	}
	return fn
}

// CreateAnimalAppointment handles the /api/create-animal-appointment endpoint and adds a row to klinikkaohjelma_kehitys.telainkaynti
func CreateAnimalAppointment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-animal-appointment/")
		decoder := json.NewDecoder(r.Body)
		var appointment structs.AnimalAppointment
		err := decoder.Decode(&appointment)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		_, err = b.Repo.CreateAnimalAppointment(appointment)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
	}
	return fn
}

func IncAppointmentIDs(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/inc-appointment/")
		ids, err := b.Repo.GetAllAppointmentIDs()
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			log.Fatal(err)
		}
		go inc(b, ids)
		json.NewEncoder(w).Encode(len(ids))
	}
	return fn
}
func inc(b *context.Base, ids []int) {
	for _, id := range ids {
		iD := strconv.Itoa(id)
		err := b.Mmqtt.Publish("modbus-event", `"appointmentId":"`+iD+`}"`)
		if err != nil {
			log.Println("inc() err: " + err.Error())
		}
	}
}

// GetAppointmentsByStates responds with an array of Appointments
// With the state numbers in path parameter states, separated by a comma.
func GetOpenAppointments(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointments/{states}")
		//var ret []structs.Appointment
		//ret = make([]structs.Appointment, 0, 0)
		date := r.URL.Query().Get("date")
		recs, err := b.Repo.GetOpenAppointments(date)
		if err != nil {
			log.Println("GetOpenAppointments error: " + err.Error())
			errRespond(err, w)
			return
		}
		recs, err = b.Repo.FillAppointments(recs)
		if err != nil {
			log.Println("GetOpenAppointments error: " + err.Error())
			errRespond(err, w)
			return
		}
		json.NewEncoder(w).Encode(recs)
	}
	return fn
}

// GetAppointmentsByStates responds with an array of Appointments
// With the state numbers in path parameter states, separated by a comma.
func GetAppointmentsByStates(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointments/{states}")
		params := mux.Vars(r)
		states := strings.Split(params["states"], ",")
		var ret []structs.Appointment
		ret = make([]structs.Appointment, 0, 0)

		for _, s := range states {
			i, err := strconv.Atoi(s)
			if errRespond(err, w) {
				return
			}
			recs, err := b.Repo.GetAppointmentsByState(i)
			if errRespond(err, w) {
				return
			}
			filledAppts, err := b.Repo.FillAppointments(recs)
			if errRespond(err, w) {
				return
			}
			ret = append(ret, filledAppts...)
		}

		json.NewEncoder(w).Encode(ret)
	}
	return fn
}

// GetAppointmentsByRange responds with a slice of Appointments, with dates
// within the range set by query parameters start and end in YYYY-MM-DD format.
func GetAppointmentsByRange(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointments/range/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.GetAppointmentsByRange(strStart, strEnd)
		if errRespond(err, w) {
			return
		}
		filledAppts, err := b.Repo.FillAppointments(res)
		if errRespond(err, w) {
			return
		}
		mapAppts := make(map[string][]structs.Appointment)
		for _, appt := range filledAppts {
			year := strconv.Itoa(appt.Date.Time.Year())
			month := strconv.Itoa(int(appt.Date.Time.Month()))
			day := strconv.Itoa(appt.Date.Time.Day())
			if len(month) == 1 {
				month = "0" + month
			}
			if len(day) == 1 {
				day = "0" + day
			}
			dateStr := year + "-" + month + "-" + day

			mapAppts[dateStr] = append(mapAppts[dateStr], appt)
		}
		json.NewEncoder(w).Encode(mapAppts)
	}
	return fn
}

// GetAppointmentsByDate responds with all Appointments for given path param date in YYYY-MM-DD format.
func GetAppointmentsByDate(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointments/date/{date}")
		params := mux.Vars(r)
		strDate := params["date"]
		if len(strDate) != 10 {
			strErrRespond("appointment date required in YYYY-MM-DD format.", w)
			return
		}
		res, err := b.Repo.GetAppointmentsByDate(strDate)
		log.Println("Found", len(res), "appointments.")
		if errRespond(err, w) {
			return
		}
		filledAppts, err := b.Repo.FillAppointments(res)
		log.Println("Filled", len(filledAppts), "appointments.")
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(filledAppts)
	}
	return fn
}

// GetAppointment responds with a single appointment for the given ID or null in case the appointment does not exist..
func GetAppointment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointment/{id}")
		params := mux.Vars(r)
		id := params["id"]
		intId, err := strconv.Atoi(id)
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetAppointment(intId)
		log.Println("GetAppointment result:", res)
		if errRespond(err, w) {
			return
		}
		appts := make([]structs.Appointment, 0)
		appts = append(appts, res)
		filledAppts, err := b.Repo.FillAppointments(appts)
		log.Println("Filled", len(filledAppts), "appointments.")
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(filledAppts[0])
	}
	return fn
}

//GetAppointmentsByAnimal returns animal appointment information by id
func GetAppointmentsByAnimal(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointments/animal/{animalid}")
		params := mux.Vars(r)
		animalID, err := strconv.Atoi(params["animalid"])
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetAppointmentsByAnimal(animalID)
		if errRespond(err, w) {
			return
		}
		filledAppts, err := b.Repo.FillAppointments(res)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(filledAppts)
	}
	return fn
}

//SearchAppointments returns all the appointments for searchterm
func SearchAppointments(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/search/appointments/{terms}")
		params := mux.Vars(r)
		terms := params["terms"]
		termArr := strings.Split(terms, ",")
		res, err := b.Repo.SearchAppointments(termArr)
		//log.Printf("Handlers SearchAppointments res %v", res)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		//json.NewEncoder(w).Encode(res)
		filledAppts, err := b.Repo.FillAppointments(res)
		if err != nil {
			log.Println("SearchAppointments error: " + err.Error())
		}
		if errRespond(err, w) {
			return
		}
		//log.Printf("Handlers SearchAppointments fileldApts %v", filledAppts)
		json.NewEncoder(w).Encode(filledAppts)
	}
	return fn
}

//GetAppointmentsByCustomer returns all customers' appointment information by customer id.
func GetAppointmentsByCustomer(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/appointments/customer/{customerid}")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["customerid"])
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetAppointmentsByCustomer(id)
		if errRespond(err, w) {
			return
		}
		log.Println("Start filling AAs to appointment")
		filledAppts, err := b.Repo.FillAppointments(res)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(filledAppts)
	}
	return fn
}

// UpdateAppointment handles the /api/update-appointment endpoint and adds a row to klinikkaohjelma_kehitys.tkaynti
func UpdateAppointment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/update-appointment/")
		decoder := json.NewDecoder(r.Body)
		var appointment structs.Appointment
		err := decoder.Decode(&appointment)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdateAppointment(appointment)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
	}
	return fn
}

//UpdateAnimalAppointment handles the /update-animal-appointment/ endpoint and adds a row to klinikkaohjelma_kehitys.telainkaynti
func UpdateAnimalAppointment(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /update-animal-appointment/")
		decoder := json.NewDecoder(r.Body)
		var appointment structs.AnimalAppointment
		err := decoder.Decode(&appointment)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdateAnimalAppointment(appointment)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
		return
	}
	return fn
}
