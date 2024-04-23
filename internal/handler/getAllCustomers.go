package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"github.com/HouseCham/customerService/internal/repository"
	"google.golang.org/grpc/codes"
)

// GetAllCustomersHandler is a handler that returns all customers from the database
func GetAllCustomersHandler() model.HttpResponse {
	log.Logger.Infoln("GetAllCustomersHandler invoked")

	// Create a channel to communicate the result back
	resultCh := make(chan model.HttpResponse)

	// Perform the operations in a goroutine
	go func() {
		defer close(resultCh)

		// Call the repository to get all customers
		customers, err := repository.GetAllCustomers()
		if err != nil {
			resultCh <- handleErrorResponse("Error getting all customers", err, codes.Internal, nil)
			return
		}

		// Send the successful response
		resultCh <- model.HttpResponse{
			StatusCode:   uint32(codes.OK),
			Data:         customers,
			HasError:     false,
			ErrorMessage: "",
		}
	}()

	// Wait for the result from the goroutine
	return <-resultCh
}