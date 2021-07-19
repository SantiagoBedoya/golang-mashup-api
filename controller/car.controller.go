package controller

import (
	"encoding/json"
	"net/http"

	"golang-mashups/service"
)

var (
	carService service.CarDetailsService
)

type controller struct{}

type CarController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

func NewCarController(service service.CarDetailsService) CarController {
	carService = service
	return &controller{}
}

func (*controller) GetCarDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := carService.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
