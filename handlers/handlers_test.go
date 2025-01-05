package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"myvet-v2-api/context"
)

var ctx context.Base
var r = httptest.NewRecorder()

// TestHandlers tests the handler package
// for now, it handles 17 handler functions from the package.
// --it can be shortened
//
// =========================================================================================
// || GetAnimalsByIDs          || GetBreedsBySpecies        || GetSpeciesByID
// || GetAppointmentsByStates  || GetAppointmentsByRange    || GetAppointmentsByDate
// || GetAppointmentsByAnimal  || SearchAppointments        || GetAppointmentsByCustomer
// || GetContactDetails        || GetCustomersByIDs         || GetWorkShiftsByRangeAndEmployee
// || GetTotalPayment          || FindErroneousPayments     || FindNonMatchingSum
// || GetPaymentsByBillnumbers || GetPayment2sByBillnumbers ||
func TestHandlers(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/contact-details/{customerids}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/customers/ids/{ids}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/work-shift/", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/appointments/{states}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/appointments/range/", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/appointments/date/{date}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/appointments/animal/{animalid}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/search/appointments/{terms}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/appointments/customer/{customerid}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/animals/ids/{ids}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/breeds/species/{speciesid}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/species/{id}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/total-payment/", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/debug1/", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/debug2/", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/payments/billnumbers/{billnumbers}", nil)
	errorCheck(err, t)
	req, err = http.NewRequest("GET", "/api/payments2/billnumbers/{billnumbers}", nil)
	errorCheck(err, t)

	res := (GetAnimalsByIDs(&ctx))
	res.ServeHTTP(r, req)
	res = (GetBreedsBySpecies(&ctx))
	res.ServeHTTP(r, req)
	res = (GetSpeciesByID(&ctx))
	res.ServeHTTP(r, req)

	res = (GetAppointmentsByStates(&ctx))
	res.ServeHTTP(r, req)
	res = (GetAppointmentsByRange(&ctx))
	res.ServeHTTP(r, req)
	res = (GetAppointmentsByDate(&ctx))
	res.ServeHTTP(r, req)
	res = (GetAppointmentsByAnimal(&ctx))
	res.ServeHTTP(r, req)
	res = (SearchAppointments(&ctx))
	res.ServeHTTP(r, req)
	res = (GetAppointmentsByCustomer(&ctx))
	res.ServeHTTP(r, req)

	res = (GetContactDetails(&ctx))
	res.ServeHTTP(r, req)

	res = (GetCustomersByIDs(&ctx))
	res.ServeHTTP(r, req)

	res = (GetWorkShiftsByRangeAndEmployee(&ctx))
	res.ServeHTTP(r, req)

	res = (GetTotalPayment(&ctx))
	res.ServeHTTP(r, req)
	res = (FindErroneousPayments(&ctx))
	res.ServeHTTP(r, req)
	res = (FindNonMatchingSum(&ctx))
	res.ServeHTTP(r, req)
	res = (GetPaymentsByBillnumbers(&ctx))
	res.ServeHTTP(r, req)
	res = (GetPayment2sByBillnumbers(&ctx))
	res.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Expected: %d \n Got: %d ", http.StatusOK, status)
	}
}

func errorCheck(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}
