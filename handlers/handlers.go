// Package classification: API Documentation
//
// Documentation(s) for all route parameters/APIs
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 3.0.0
//	   Contact: Ozan <ozan@adaptek.fi> https://adaptek.fi
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handlers

import (
	"encoding/json"

	"log"
	"net/http"

	"myvet-v2-api/context"
	"myvet-v2-api/mws"
	"myvet-v2-api/structs"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store *sessions.CookieStore

/**
 * Middleware functions
 * TODO: Create a middleware.go file and move these there
 */

// checkLogin only checks if the user is either logged in OR hitting the /api/login/ endpoint
// wanting to log in.
// Session is stored server side, and absolutely no data is transported to the client in the cookies.
func checkLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ses, err := store.Get(r, "session-name")
		if err != nil {
			log.Println("checkLogin err: " + err.Error())
		}
		log.Println("Is new?", ses.IsNew)
		user := ses.Values["user"]
		log.Println("user ", user)
		log.Println(r.URL.String())
		if user != nil && user != "" || (r.URL.String() == "/api/login/" && r.Method == "POST") {
			log.Println("Authenticated/authenticating user ", user)
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}

// Sets the default response content type to be application/json
func ServeJSONMW(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
	return fn
}

//InitRoutes initializes all route parameters.
func InitRoutes(c *context.Base) error {
	log.Println("Initializing routing.")

	c.Router = mux.NewRouter()
	// TODO: MAYBE MOVE TO CONFIG.
	store = sessions.NewCookieStore([]byte("secret"))

	//Add middleware

	//authentication
	if !c.Config.NoAuth {
		c.Router.Use(checkLogin)
	}

	//Default content-type
	c.Router.Use(ServeJSONMW)
	//logging
	//c.Router.Use(c.LogRequest)

	// Legacy endpoint due to removal asap, but waiting for security measures.
	c.Router.HandleFunc("/oauth/check_token", CheckToken(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/login/", Login(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/logout/", Logout(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/inc-appointment/", IncAppointmentIDs(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/receivables/date/{date}", GetReceivablesByDate(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/visits/worklist/{date}", GetWorkList(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/visits/reservation/", GetAllReservations(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/visits/receivable/", GetAllReceivables(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointments/open/", GetOpenAppointments(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointments/states/{states}", GetAppointmentsByStates(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointments/date/{date}", GetAppointmentsByDate(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointments/range/", GetAppointmentsByRange(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointments/customer/{customerid}", GetAppointmentsByCustomer(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointment/{id}", GetAppointment(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/appointments/animal/{animalid}", GetAppointmentsByAnimal(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments/debug/", FindErroneousPayments(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments/debug2/", FindNonMatchingSum(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments/billnumbers/{billnumbers}", GetPaymentsByBillnumbers(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments2/billnumbers/{billnumbers}", GetPayment2sByBillnumbers(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/total-payment/", GetTotalPayment(c)).Methods("GET")

	// c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments/dates/", GetPaymentsByDates(c)).Methods("POST")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/customers/id/", GetCustomersByIDs(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/customer/{customerId}/animals", GetCustomerAnimals(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/work-shift/", GetWorkShiftsByRangeAndEmployee(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/shifts/", GetWorkShiftsByRange(c)).Methods("GET")

	// Employees are all fetched at once on the front-end.
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/employees/", GetEmployees(c)).Methods("GET")

	// Medicines are all fetched at once on the front-end.
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/medicines/", GetMedicines(c)).Methods("GET")

	// Units are all fetched at once on the front-end.
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/units/", GetUnits(c)).Methods("GET")
	// Species are all fetched at once on the front-end.
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/species/", GetAllSpecies(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/species/{id}", GetSpeciesByID(c)).Methods("GET")

	// 'api/breeds/species/' + speciesId
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/breeds/", GetAllBreeds(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/breeds/species/{speciesid}", GetBreedsBySpecies(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/customers/", GetAllCustomers(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/customer/{id}", GetCustomer(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/customers/ids/{ids}", GetCustomersByIDs(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/contact-details/{customerids}", GetContactDetails(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-employee/", CreateEmployee(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-employee/", UpdateEmployee(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-payment/", CreatePayment(c)).Methods("POST")
	// tsuoritus2 payments used
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-payments/", CreatePayment2s(c)).Methods("POST")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-payment/", UpdatePayment(c)).Methods("POST")
	//tapiola reply
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/tapiola-reply/", CreateTapiolaReply(c)).Methods("POST")
	// tapiola communication
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/tapiola-communication/", CreateTapiolaCommmunication(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-appointment/", CreateAppointment(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-appointment/", UpdateAppointment(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-customer/", CreateCustomer(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-customer/", UpdateCustomer(c)).Methods("POST")
	// creates a new animal
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/animals/", CreateAnimal(c)).Methods("POST")
	// updates an existing animal
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/animals/{animalID}", UpdateAnimal(c)).Methods("POST")
	//c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-animal/{animalID}", UpdateAnimal(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-animal-appointment/", CreateAnimalAppointment(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-animal-appointment/", UpdateAnimalAppointment(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/create-contact/", CreateContact(c)).Methods("POST")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/update-contact/", UpdateContact(c)).Methods("POST")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/animal/{id}", GetAnimal(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/animals/ids/{ids}", GetAnimalsByIDs(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/temporary-page/{id}", GetTemporaryPage(c)).Methods("GET")

	/////////////////////////////////////////////////////////////////
	// The API endpoints implementing report data start from here.

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/daily-income/{date}", GetDailyIncomeReport(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/daily-total/{date}", GetDailyTotal(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/monthly-income/", GetMonthlyIncomeReport(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/paid-receivables/{date}", GetPaidReceivables(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/paid-receivables/", GetPaidReceivables(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/range-total/", GetRangeTotal(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/medicine/", GetMedicineReports(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/medicine/count/", GetMedicineReportsCount(c)).Methods("GET")
	//c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/medicine/count/", GetMedicineReportsCount(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/annual-sales/{year}", GetAnnualSales(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/reports/daily-sales/{date}", GetDailySales(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/vat-rates/", GetVATRates(c)).Methods("GET")
	/////////////////////////////////////////////////////////////////
	// The API endpoints implementing data search start from here.
	// The searches can be complex and therefor are possibly time consuming and hence they are separated
	// from normal database getters, although at the moment implemented as SQL queries.

	// api/search/appointments/' + termsStr
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/search/appointments/{terms}", SearchAppointments(c)).Methods("GET")
	// api/search/customersanimals/' + termsStr
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/search/customersAnimals/{terms}", SearchCustomersAnimals(c)).Methods("GET")
	// api/search/customers/{terms}
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/search/customers/{terms}", SearchCustomers(c)).Methods("GET")

	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/purposes/", GetAllPurposes(c)).Methods("GET")

	// weight history
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/weight-history/{animalID}", GetWeightHistory(c)).Methods("GET")
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/weight-history/{animalID}/new", GetWeightHistory(c)).Methods("POST")

	getDoc := c.Router.Methods(http.MethodGet).Subrouter()
	opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	sh := middleware.Redoc(opts, nil)

	getDoc.Handle("/docs", sh)
	getDoc.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	log.Println("Initializing websocket route " + c.Config.APIPrefix + "/" + c.Config.WebsocketRoute)
	c.WS = mws.NewHub(WSMsgRcvd(c))
	go c.WS.Run()
	c.Router.HandleFunc("/"+c.Config.APIPrefix+"/"+c.Config.WebsocketRoute, getServeWS(c))

	//test for logging
	//logging
	//c.Router.Use(c.LogRequest)

	///////////////
	// Catch-all route to log what the frontend does.
	//////////
	c.Router.PathPrefix("/").HandlerFunc(DefaultHandler())

	/////////////////////////////////////////
	// DEBUG endpoints, implement these last,
	// or be more specific, when something really doesn't work :).

	// api/receivables/debug/
	// This takes POST object, { start: string, end: string } start and end strings contain date in YYYY-MM-DD fornat.
	// c.Router.HandleFunc("/"+c.Config.APIPrefix+"/receivables/debug/", GetReceivablesDebug(c)).Methods("POST")

	// api/payments/debug
	// This takes POST object, { start: string, end: string } start and end strings contain date in YYYY-MM-DD fornat.
	// c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments/debug/", GetPaymentsDebug(c)).Methods("POST")

	// api/payments/debug2
	// This takes POST object, { start: string, end: string } start and end strings contain date in YYYY-MM-DD fornat.
	// c.Router.HandleFunc("/"+c.Config.APIPrefix+"/payments/debug2/", GetPaymentsDebug2(c)).Methods("POST")
	return nil

}

func errRespond(err error, w http.ResponseWriter) bool {
	if err != nil {
		json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
		return true
	}
	return false
}

func strErrRespond(msg string, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(structs.JSONErr{Error: msg})
}

// WSMsgRcvd handles logic of receiving websocket messages.
func WSMsgRcvd(c *context.Base) func(clientID string, message string) {
	fn := func(clientID string, message string) {
		var wsMsg Msg
		err := json.Unmarshal([]byte(message), &wsMsg)
		if err != nil {
			log.Println(err)
		}
		if wsMsg.Type == "MESSAGE" {
			for m := range c.WS.Clients {
				if m.ID == clientID {
					wsResponse, err := json.Marshal(Msg{Type: "TEMP-MESSAGE", Payload: "", Target: ""})
					if err != nil {
						log.Println(err)
					}
					m.Send <- wsResponse
				}
			}
			return
		}
		if wsMsg.Type == "User Message" {
			log.Println("Direct message from user to:", wsMsg.Target)
		}
		log.Println("Received websocket message from: " + clientID)
		log.Println("Message content: " + message)
		log.Println(wsMsg.Type)
		log.Println(wsMsg.Target)

	}
	return fn
}

func getServeWS(c *context.Base) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Run getServeWS")

	}
	return fn
}

func DefaultHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Default route hit:")
		params := mux.Vars(r)
		log.Println(params)
		log.Println(r.URL)
		log.Println(r.Form)
		log.Println(r.Header)

		json.NewEncoder(w).Encode("Hit non-existant end-point: " + r.URL.String() + " method: " + r.Method)
		log.Println("töttöröö")
	}
	return fn
}

//Credentials holds user credential tokens
type Credentials struct {
	User        string `json:"user"`
	Pass        string `json:"pass"`
	EncodedPass string `json:"encodedPass"`
}

func Login(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("/api/login/")
		decoder := json.NewDecoder(r.Body)
		log.Println(r.Body)
		var creds Credentials
		err := decoder.Decode(&creds)
		if err != nil {
			log.Println(err)
			log.Println("skipping blank credentials")
		}
		log.Println("User", creds.User, "logging in.")
		ses, err := store.Get(r, "session-name")
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("täs ois", ses.Values["user"])
		log.Println("/api/login/ user: ", ses.Values["user"])
		ses.Values["user"] = creds.User
		log.Println("Getting user", creds.User, "from the database.")
		if creds.User == "" {
			strErrRespond("No username given", w)
			return
		}

		user, err := b.Repo.GetUser(creds.User)

		log.Println(err)
		if creds.Pass == "" && creds.EncodedPass == "" {
			strErrRespond("Neither pass or encodedPass were included in the login request:"+r.URL.RawQuery, w)
			return
		}
		if creds.Pass != "" {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Pass))
			if errRespond(err, w) {
				return
			}
		} else {
			if creds.EncodedPass != user.Password {
				strErrRespond("Encoded password did not match.", w)
				return
			}
		}

		log.Println(user)
		log.Println("before saving ses")
		log.Println(ses.Values["user"])
		err = ses.Save(r, w)
		if err != nil {
			log.Println("Login error: " + err.Error())
		}

		var userDto structs.UserDTO
		userDto.Enabled = user.Enabled
		userDto.Expires = user.Expires
		userDto.Lang = user.Lang
		userDto.Locked = user.Locked
		userDto.Permissions = user.Permissions
		userDto.Picnum = user.Picnum
		userDto.Role = user.Role
		userDto.UserName = user.UserName
		log.Println(ses.Values["user"])
		// err = ses.Save(r, w)
		err = store.Save(r, w, ses)
		if err != nil {
			log.Println("Login error: " + err.Error())
		}
		json.NewEncoder(w).Encode(userDto)
		// check(json.NewEncoder(w).Encode("User " + fmt.Sprintf("%v", ses.Values["user"]) + " logged in."))
	}
	return fn
}
func Logout(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("handlers.go.Logout()")
		ses, err := store.Get(r, "session-name")
		if err != nil {
			log.Println("Logout error: " + err.Error())
			return
		}
		log.Println("Logging out", ses.Values["user"])
		ses.Values["user"] = ""
		ses.Save(r, w)
		json.NewEncoder(w).Encode("Logged out.")
	}
	return fn
}

// GetMedicines responds with a list of all medicines in the database..
func GetMedicines(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		res, err := b.Repo.GetMedicines()
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
	}
	return fn
}

// GetMEdicines responds with a list of all medicines in the database..
func GetAllPurposes(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		res, err := b.Repo.GetAllPurposes()
		if errRespond(err, w) {
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

// GetUnits responds with a list of all units in the database.
// In this case "unit" refers to unit of measures for eg. medicines or feeds.
func GetUnits(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		res, err := b.Repo.GetUnits()
		if err != nil {
			json.NewEncoder(w).Encode(structs.JSONErr{Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	return fn
}

// CheckToken is a mock endpoint required for the system upgrade.
func CheckToken(b *context.Base) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /oauth/check_token")
		json.NewEncoder(w).Encode("TOKEN REPLY!")
	}
	return fn
}

//Entries contains required fields for making handler-functions.
type Entries struct {
	IDs    []string `json:"ids"`
	Dates  []string `json:"dates"`
	States []string `json:"states"`
}

// Msg is a generic message structure.
type Msg struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Target  string `json:"target"`
}
