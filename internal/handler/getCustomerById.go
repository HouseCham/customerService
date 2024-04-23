package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"github.com/HouseCham/customerService/internal/repository"
	"google.golang.org/grpc/codes"
)

// GetCustomerByIdHandler is a handler that receives a customer ID and returns the customer from the database
func GetCustomerByIdHandler(customerId uint) model.HttpResponse {
	log.Logger.Infoln("GetCustomerByIdHandler invoked")

	// Create a channel to communicate the result back
	resultCh := make(chan model.HttpResponse)

	// Perform the operations in a goroutine
	go func() {
		defer close(resultCh)

		// Call the repository to get the customer
		customer, err := repository.GetCustomerByID(customerId)
		if err != nil {
			resultCh <- handleErrorResponse("Error getting customer", err, codes.Internal, nil)
			return
		}

		// Send the successful response
		resultCh <- model.HttpResponse{
			StatusCode:   uint32(codes.OK),
			Data:         customer,
			HasError:     false,
			ErrorMessage: "",
		}
	}()

	// Wait for the result from the goroutine
	return <-resultCh
}