package handlers

import (
	"encoding/json"
	"log"
	"myvet-v2-api/context"
	"myvet-v2-api/structs"
	"net/http"
	"strconv"
)

// CreateEmployee handles the /api/create-employee endpoint and adds a row to klinikkaohjelma_kehitys.ttyontekija
func CreateEmployee(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/create-employees/")
		decoder := json.NewDecoder(r.Body)
		var employee structs.Employee
		err := decoder.Decode(&employee)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}

		lastID, err := b.Repo.CreateEmployee(employee)
		if errRespond(err, w) {
			return
		}
		employee.EmployeeID = lastID
		json.NewEncoder(w).Encode(employee)
		return
	}
	return fn
}

//UpdateEmployee handles the /api/update-employee/ endpoint and adds a row to klinikkaohjelma_kehitys.ttyontekija
func UpdateEmployee(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/update-employee/")
		decoder := json.NewDecoder(r.Body)
		var employee structs.Employee
		err := decoder.Decode(&employee)
		if errRespond(err, w) {
			return
		}
		rowsAff, err := b.Repo.UpdateEmployee(employee)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode("Updated " + strconv.Itoa(rowsAff) + " rows.")
		return
	}
	return fn
}

// GetEmployees responds with a list of all clinic Employees.
func GetEmployees(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		res, err := b.Repo.GetEmployees()
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		log.Println("GetEmployees()")
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

// GetWorkShiftsByRangeAndEmployee returns list of work shift information given by date.
func GetWorkShiftsByRangeAndEmployee(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/work-shift/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		strID := r.URL.Query().Get("employeeid")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.GetWorkShiftsByRangeAndEmployee(strStart, strEnd, strID)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

// GetWorkShiftsByRange returns list of work shift information given by date.
func GetWorkShiftsByRange(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/work-shift/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.GetWorkShiftsByRange(strStart, strEnd)
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}
