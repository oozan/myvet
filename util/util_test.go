package util

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterBusinessDays(t *testing.T) {
	t1 := time.Now()
	// assure that error is given when called with negative "days" argument
	var err error
	_, err = AfterBusinessDays(t1, -1)
	if err == nil {
		t.Errorf("Expecting error when calling AfterBusinessDays with negative 'days' argument")
	}
	// assure that that time isn't changed when calling with 0 days argument
	var res time.Time
	res, err = AfterBusinessDays(t1, 0)
	if err != nil {
		t.Errorf("No error expected when calling AfterBusinessDays(t1,0)")
	}
	fmt.Println(res)
	if !t1.Equal(res) {
		t.Errorf("Expecting resulting date to be the same as the date argument when calling AfterBusinessDays(date,0)")
	}
	// Tuesday 30.6.2020 - expect the result to be thursday 2.7.2020
	// after calling AfterBusinessDays with days = 2
	//then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fiT, _ := time.LoadLocation("Europe/Helsinki")
	start := time.Date(2020, 6, 30, 0, 0, 0, 0, fiT)
	expected := time.Date(2020, 7, 2, 0, 0, 0, 0, fiT)
	res, err = AfterBusinessDays(start, 2)
	fmt.Println(start)
	fmt.Println(expected)
	if err != nil {
		t.Errorf("error: %s\n", err.Error())
	}
	if !res.Equal(expected) {
		t.Errorf("Expecting %s but got %s\n", expected.String(), res.String())
	}
	//friday
	start = time.Date(2020, 7, 3, 0, 0, 0, 0, fiT)
	//monday
	expected = time.Date(2020, 7, 6, 0, 0, 0, 0, fiT)
	res, err = AfterBusinessDays(start, 1)
	fmt.Println(res)
	fmt.Println(start)
	fmt.Println(expected)
	if err != nil {
		t.Errorf("error: %s\n", err.Error())
	}
	if !res.Equal(expected) {
		t.Errorf("Expecting %s but got %s\n", expected.String(), res.String())
	}
}
