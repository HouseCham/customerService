package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"github.com/HouseCham/customerService/internal/validator"
	"github.com/HouseCham/customerService/internal/repository"
	"google.golang.org/grpc/codes"
)

// UpdateCustomerHandler is a handler that receives a customer and updates it in the database
func UpdateCustomerHandler(requestBody *model.Customer) model.HttpResponse {
	log.Logger.Infoln("updateCustomerHandler invoked")

	// Create a channel to communicate the result back
	resultCh := make(chan model.HttpResponse)

	// Perform the operations in a goroutine
	go func() {
		defer close(resultCh)

		// Check if the customer exists
		if _, err := repository.GetCustomerByID(requestBody.ID); err != nil {
			resultCh <- handleErrorResponse("Customer does not exist", err, codes.NotFound, nil)
			return
		}

		// Validate the request body
		if err := validator.Validate.Struct(requestBody); err != nil {
			resultCh <- handleErrorResponse("Error validating request body", err, codes.InvalidArgument, validator.GetValidatorErrorMessage(err))
			return
		}

		// Call the repository to update the customer
		err := repository.UpdateCustomerDB(requestBody)
		if err != nil {
			resultCh <- handleErrorResponse("Error updating customer", err, codes.Internal, nil)
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