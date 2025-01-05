package handlers

import (
	"encoding/json"
	"log"
	"myvet-v2-api/context"
	"myvet-v2-api/structs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetDailyIncomeReport(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/daily-income/{date}")
		params := mux.Vars(r)
		strDate := params["date"]
		if len(strDate) != 10 {
			strErrRespond("report date required in YYYY-MM-DD format.", w)
			return
		}
		res, err := b.Repo.GetDailyIncomeReport(strDate)
		log.Println("Found", len(res), "daily income report items.")
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

func GetMonthlyIncomeReport(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/monthly-income/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.GetMonthlyIncomeReport(strStart, strEnd)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

func GetPaidReceivables(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		strDate := params["date"]
		if strDate == "" {
			log.Println("Handling /api/reports/paid-receivables/")
			strStart := r.URL.Query().Get("start")
			strEnd := r.URL.Query().Get("end")
			log.Println("Range from", strStart, "to", strEnd)
			if len(strStart) != 10 || len(strEnd) != 10 {
				json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
				return
			}
			res, err := b.Repo.GetMonthlyPaidReceivables(strStart, strEnd)
			log.Println("Found", len(res), "monthly paid receivables report items.")
			if errRespond(err, w) {
				return
			}
			json.NewEncoder(w).Encode(res)
			return
		} else {
			log.Println("Handling /api/reports/paid-receivables/" + strDate)
			if len(strDate) != 10 {
				strErrRespond("report date required in YYYY-MM-DD format.", w)
				return
			}
			res, err := b.Repo.GetDailyPaidReceivables(strDate)
			log.Println("Found", len(res), "daily paid receivables report items.")
			if errRespond(err, w) {
				return
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	return fn
}

func GetAnnualSales(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/annual-sales/{year}")
		params := mux.Vars(r)
		year := params["year"]
		vetStr := r.URL.Query().Get("vet")
		log.Println("strVet:", vetStr)
		vetInt, err := strconv.Atoi(vetStr)
		log.Println("vetint", vetInt)
		var vet *int
		if err == nil {
			vet = &vetInt
		}
		log.Println("Vet is ", vet)
		if len(year) != 4 {
			strErrRespond("year required.", w)
			return
		}
		i, err := strconv.Atoi(year)
		if errRespond(err, w) {
			return
		}
		res, err := b.Repo.GetAnnualSales(i, vet)
		if errRespond(err, w) {
			return
		}
		log.Println("Found", len(res), "annual sales reports.")
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

func GetDailyTotal(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/daily-income/{date}")
		params := mux.Vars(r)
		strDate := params["date"]
		if len(strDate) != 10 {
			strErrRespond("report date required in YYYY-MM-DD format.", w)
			return
		}
		res, err := b.Repo.GetDailyTotal(strDate)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

func GetRangeTotal(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/range-total/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd})
			return
		}
		res, err := b.Repo.GetRangeTotal(strStart, strEnd)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

func GetMedicineReports(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/medicine/")
		strStart := r.URL.Query().Get("start")
		strEnd := r.URL.Query().Get("end")
		strName := r.URL.Query().Get("name")
		if len(strStart) != 10 || len(strEnd) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + strStart + " and " + strEnd + " and name: " + strName})
			return
		}
		res, err := b.Repo.GetMedicineReports(strStart, strEnd, strName)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

/*
 * handler function for api/reports/medicine/count/
 */
func GetMedicineReportsCount(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")
		name := r.URL.Query().Get("name")
		//start := r.URL.Query().Get("start")
		log.Println("start = " + start)
		//end := r.URL.Query().Get("end")
		log.Println("end = " + end)
		//name := r.URL.Query().Get("name")
		if len(start) != 10 || len(end) != 10 {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: "Start and end query parameters required in YYYY-MM-DD format, got " + start + " and " + end + " and name: " + name})
			return
		}
		res, err := b.Repo.GetMedicineReportsCount(start, end, name)
		if err != nil {
			log.Printf("err=%s\n", err.Error())
		}
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

func GetDailySales(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/reports/daily-sales/{date}")
		var vet *int = nil
		params := mux.Vars(r)
		strDate := params["date"]
		if len(strDate) != 10 {
			strErrRespond("report date required in YYYY-MM-DD format.", w)
			return
		}
		vetStr := r.URL.Query().Get("vet")
		if vetStr != "" {
			log.Println("strVet:", vetStr)
			vetInt, errConv := strconv.Atoi(vetStr)
			if errRespond(errConv, w) {
				return
			}
			log.Println("vetint", vetInt)
			vet = &vetInt
		}
		if vet != nil {
			log.Println("*vet: ", *vet)
		}
		res, err := b.Repo.GetDailySales(strDate, vet)
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}
