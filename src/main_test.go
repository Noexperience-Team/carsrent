package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	model "github.com/Noexperience-Team/carsrent/src/Models"
	db "github.com/Noexperience-Team/carsrent/src/Repository"
	router "github.com/Noexperience-Team/carsrent/src/Routers"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	route := router.NewRouter()
	route.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetCars(t *testing.T) {
	db.Connect("../config/config.prod.yml")
	t.Parallel()

	res := new(model.Response)

	req, _ := http.NewRequest("GET", "/api/cars", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var decode = json.NewDecoder(response.Body)
	err := decode.Decode(res)
	if err != nil {
		t.Errorf("TestGetCars:Expected an empty array. Got %s", err)
	}
	if body := response.Body.String(); res.Data == nil {
		t.Errorf("TestGetCars:Expected an empty array. Got %s", body)

	} else {
		t.Logf("TestGetCars:%v", body)

	}
}

func TestAddCar(t *testing.T) {
	db.Connect("../config/config.prod.yml")
	t.Parallel()

	payload := []byte(`{
		"model": "Tesla M3",
		"registration": "BTS811112",
		"mileage":6003
	}`)
	res := new(model.Response)
	req, _ := http.NewRequest("POST", "/api/cars", bytes.NewBuffer(payload))
	response := executeRequest(req)
	var decode = json.NewDecoder(response.Body)
	err := decode.Decode(res)
	if err != nil {
		t.Errorf("TestRent:Expected an empty array. Got %s", err)
	}
	if body := res.Data; res.Error != "" {
		t.Errorf("TestRent:Got  an error :%s", res.Error)

	} else {
		t.Logf("TestRent:%v", body)

	}

}
func TestRent(t *testing.T) {
	db.Connect("../config/config.prod.yml")
	t.Parallel()

	res := new(model.Response)
	req, _ := http.NewRequest("POST", "/api/cars/BTS811112/rentals", nil)
	response := executeRequest(req)
	var decode = json.NewDecoder(response.Body)
	err := decode.Decode(res)
	if err != nil {
		t.Errorf("TestAddCar:Expected an empty array. Got %s", err)
	}
	if body := res.Data; res.Error != "" {
		t.Errorf("TestAddCar:Got  an error :%s", res.Error)

	} else {
		t.Logf("TestAddCar:%v", body)

	}

}
func TestReturnsCar(t *testing.T) {
	db.Connect("../config/config.prod.yml")
	t.Parallel()

	form := url.Values{}
	form.Add("mileage", "1000")

	res := new(model.Response)
	req, _ := http.NewRequest("POST", "/api/cars/BTS811112/returns", strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)

	var decode = json.NewDecoder(response.Body)
	err := decode.Decode(res)

	if err != nil {
		t.Errorf("TestReturnsCar:parsing response error: %s", err)
	}
	if body := res.Data; res.Error != "" {
		t.Errorf("TestReturnsCar: Got  an error :%v", res.Error)

	} else {
		t.Logf("TestReturnsCar:%v", body)

	}

}
