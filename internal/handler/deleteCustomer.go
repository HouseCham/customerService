package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"github.com/HouseCham/customerService/internal/repository"
	"google.golang.org/grpc/codes"
)

// DeleteCustomerHandler is a handler that receives a customer ID and deletes it from the database
func DeleteCustomerHandler(customerId uint) model.HttpResponse {
	log.Logger.Infoln("DeleteCustomerHandler invoked")

	// Create a channel to communicate the result back
	resultCh := make(chan model.HttpResponse)

	// Perform the operations in a goroutine
	go func() {
		defer close(resultCh)

		// Check if the customer exists
		_, err := repository.GetCustomerByID(customerId)
		if err != nil {
			resultCh <- handleErrorResponse("Customer not found", err, codes.NotFound, nil)
			return
		}

		// Call the repository to delete the customer
		err = repository.DeleteCustomerDB(
			&model.Customer{
				ID: customerId,
			},
		)
		if err != nil {
			resultCh <- handleErrorResponse("Error deleting customer", err, codes.Internal, nil)
			return
		}
		// Send the successful response
		resultCh <- model.HttpResponse{
			StatusCode:   uint32(codes.OK),
			Data:         nil,
			HasError:     false,
			ErrorMessage: "",
		}
	}()

	// Wait for the result from the goroutine
	return <-resultCh
}