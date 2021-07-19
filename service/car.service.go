package service

import (
	"fmt"
	"log"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s\n", carServiceUrl)

	// Call the external api
	resp, err := client.Get(carServiceUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Write response to channel (TODO)
	carDataChannel <- resp
}
