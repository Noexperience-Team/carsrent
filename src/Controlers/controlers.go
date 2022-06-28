package Controlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/Noexperience-Team/carsrent/src/Models"
	db "github.com/Noexperience-Team/carsrent/src/Repository"
	"github.com/gorilla/mux"
)

var Db db.Db

func GetCars(w http.ResponseWriter, r *http.Request) {
	Db.DB = db.DB
	response := new(model.Response)
	res := new([]model.Car)
	err := Db.GetCars(res)
	if err.Error != nil {

		w.WriteHeader(http.StatusBadRequest)

		response.Error = err.Error.Error()
		json.NewEncoder(w).Encode(response)
		return

	}
	response.Data = res
	json.NewEncoder(w).Encode(response)
}
func AddCars(w http.ResponseWriter, r *http.Request) {
	response := new(model.Response)
	Db.DB = db.DB
	req := new(model.Car)
	var decode = json.NewDecoder(r.Body)
	err := decode.Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return

	}
	res := Db.FindCar(req)

	if req.Id > 0 {
		w.WriteHeader(http.StatusFound)
		response.Error = " the car is already exist"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		res = Db.AddCar(req)

		if res.Error != nil {

			w.WriteHeader(http.StatusBadRequest)
			response.Error = res.Error.Error()
			json.NewEncoder(w).Encode(response)
			return
		}
		response.Data = req
		json.NewEncoder(w).Encode(response)
	}

}
func UpdateCarRentStatus(w http.ResponseWriter, r *http.Request) {
	Db.DB = db.DB
	response := new(model.Response)
	req := new(model.Car)
	req.Registration = mux.Vars(r)["registration"]
	Db.FindCar(req)
	if mux.Vars(r)["rent"] == "returns" {
		if req.Rented == false {
			w.WriteHeader(http.StatusFound)
			response.Error = " the car is already not rented"
			json.NewEncoder(w).Encode(response)
			return
		} else {
			r.ParseForm()
			Mileage, err := strconv.Atoi(r.FormValue("mileage"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				response.Error = err.Error()
				json.NewEncoder(w).Encode(response)
				return

			}
			req.Mileage = req.Mileage + Mileage
			req.Rented = false
		}

	} else if mux.Vars(r)["rent"] == "rentals" {
		if req.Rented == true {
			w.WriteHeader(http.StatusFound)
			response.Error = " the car is already rented"
			json.NewEncoder(w).Encode(response)
			return
		} else {
			req.Rented = true
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "param rent status  is not allowed you should use rentals or returns"
		json.NewEncoder(w).Encode(response)
		return
	}
	res := Db.UpdateCar(req)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = res.Error.Error()
		json.NewEncoder(w).Encode(response)
		return
	} else if res.RowsAffected < 1 {

		if req.Id < 1 {
			w.WriteHeader(http.StatusNotFound)
			response.Error = "no car has been found"

			json.NewEncoder(w).Encode(response)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			response.Error = " the rent status is already updated"
			json.NewEncoder(w).Encode(response)
		}

		return
	}
	response.Data = *req
	json.NewEncoder(w).Encode(response)
}
