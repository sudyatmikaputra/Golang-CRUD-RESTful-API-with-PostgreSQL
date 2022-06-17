package router

import (
	"gotest/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/car", controller.GetAllCars).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/car", controller.AddCar).Methods("POST", "OPTIONS")
	
	router.HandleFunc("/api/car/{id}", controller.GetCar).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/car/{id}", controller.UpdateCar).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/car/{id}", controller.DeleteCar).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/carbrand/{brand}", controller.GetCarByBrand).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/carbrand/{brand}", controller.DeleteCarByBrand).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/carname/{name}", controller.GetCarByName).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/carname/{name}", controller.UpdateCarByName).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/carname/{name}", controller.DeleteCarByName).Methods("DELETE", "OPTIONS")

	return router
}