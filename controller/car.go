package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"log"
	"net/http"
	"gotest/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      	int64  	`json:"id,omitempty"`
	Message 	string 	`json:"message,omitempty"`
	CarBrand	string	`json:"car_brand,omitempty"`
	CarName		string	`json:"car_name,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Car  `json:"data"`
}

// CRUD BY ID

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatalf("Fail to decode body request.  %v", err)
	}
	insertID := models.AddCar(car)
	res := response{
		ID:      insertID,
		Message: "Car has been added",
	}
	json.NewEncoder(w).Encode(res)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Fail to convert String to INT.  %v", err)
	}

	car, err := models.GetCar(int64(id))
	if err != nil {
		log.Fatalf("Fail to get book data. %v", err)
	}
	json.NewEncoder(w).Encode(car)
}


func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cars, err := models.GetAllCars()
	if err != nil {
		log.Fatalf("Fail to get data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = cars

	json.NewEncoder(w).Encode(response)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Fail to convert String to INT.  %v", err)
	}

	var car models.Car

	err = json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatalf("Fail to decode body request.  %v", err)
	}

	updatedRows := models.UpdateCar(int64(id), car)

	msg := fmt.Sprintf("Car list successfully updated. Total updated data is %v rows/record", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Fail to convert String to INT. %v", err)
	}

	deletedRows := models.DeleteCar(int64(id))

	msg := fmt.Sprintf("Car successfully deleted. Total deleted data is %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

//-------------------------------------------------------------------------------------------------------------//

// CRUD BY BRAND

func GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	brand := params["brand"]

	cars, err := models.GetCarByBrand(string(brand))
	if err != nil {
		log.Fatalf("Fail to get book data. %v", err)
	}
	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = cars
	
	json.NewEncoder(w).Encode(response)
}


func DeleteCarByBrand(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	brand := params["brand"]

	deletedRows := models.DeleteCarByBrand(string(brand))
	msg := fmt.Sprintf("Car successfully deleted. Total deleted data is %v", deletedRows)
	res := response{
		CarBrand: string(brand),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}


//-------------------------------------------------------------------------------------------------------------//

// CRUD BY NAME

func GetCarByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	carname := params["name"]

	car, err := models.GetCarByName(string(carname))
	if err != nil {
		log.Fatalf("Fail to get book data. %v", err)
	}
	json.NewEncoder(w).Encode(car)
}


func UpdateCarByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carname := params["name"]

	var car models.Car

	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatalf("Fail to decode body request.  %v", err)
	}

	updatedRows := models.UpdateCarByName(string(carname), car)
	
	msg := fmt.Sprintf("Car list successfully updated. Total updated data is %v rows/record", updatedRows)
	res := response{
		CarName: string(carname),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteCarByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carname := params["name"]
	deletedRows := models.DeleteCarByName(string(carname))

	msg := fmt.Sprintf("Car successfully deleted. Total deleted data is %v", deletedRows)
	res := response{
		CarName: string(carname),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}