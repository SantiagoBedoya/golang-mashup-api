package service

import (
	"fmt"
	"log"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s\n", ownerServiceUrl)

	// Call the external api
	resp, err := client.Get(ownerServiceUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Write response to channel (TODO)
	ownerDataChannel <- resp
}
