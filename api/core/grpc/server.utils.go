package grpc

import (
	"github.com/HouseCham/customerService/internal/model"
	"google.golang.org/grpc/codes"
)

// Function to map InsertCustomerCommand to InsertCustomerModel
func mapInsertCustomerRequestToModel(request *InsertCustomerCommand) *model.Customer {
	return &model.Customer{
		FirstName:   request.FirstName,
		SecondName:  request.SecondName,
		LastNameP:   request.LastNameP,
		LastNameM:   request.LastNameM,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    request.Password,
	}
}

// transformValidatorErrors transforms an array of ValidatorError to an array of ValidationError
func transformValidatorErrors(validatorErrors []model.ValidatorError) []*ValidationError {
	var validationErrors []*ValidationError
	for _, ve := range validatorErrors {
		validationErrors = append(validationErrors, &ValidationError{
			Tag:   ve.Tag,
			Field: ve.Field,
			Err:   ve.Err,
		})
	}
	return validationErrors
}

// InsertCustomerResponse is the response struct for the InsertCustomer gRPC endpoint
func handleResponseErrors(handlerResponse model.HttpResponse) *InsertCustomerResponse {
	response := &InsertCustomerResponse{}
	// Check if there was an error
	if handlerResponse.HasError {
		response.HasError = true
		response.ErrorMessage = handlerResponse.ErrorMessage

		// Handle validation errors
		if handlerResponse.StatusCode == uint32(codes.InvalidArgument) {
			response.ErrorMessage = "invalid request body"
			response.ValidationErrors = transformValidatorErrors(handlerResponse.Data.([]model.ValidatorError))
		}
	} else {
		response.HasError = false
		response.NewId = uint32(handlerResponse.Data.(uint))
		response.StatusCode = handlerResponse.StatusCode
	}
	return response
}
