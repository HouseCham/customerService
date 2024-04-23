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

// Function to map UpdateCustomerCommand to Customer model
func mapUpdateCustomerRequestToModel(request *UpdateCustomerCommand) *model.Customer {
	return &model.Customer{
		FirstName:   request.FirstName,
		SecondName:  request.SecondName,
		LastNameP:   request.LastNameP,
		LastNameM:   request.LastNameM,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
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
func handleInsertResponseErrors(handlerResponse model.HttpResponse) *InsertCustomerResponse {
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

// UpdateCustomerResponse is the response struct for the UpdateCustomer gRPC endpoint
func handleUpdateResponseErrors(handlerResponse model.HttpResponse) *UpdateCustomerResponse {
	response := &UpdateCustomerResponse{}
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
		response.StatusCode = handlerResponse.StatusCode
	}
	return response
}
// handleGetResponseErrors is a function that handles the response errors for the GetCustomer gRPC endpoint
func handleGetResponseErrors(handlerResponse model.HttpResponse) *GetCustomerResponse {
	response := &GetCustomerResponse{}
	// Check if there was an error
	if handlerResponse.HasError {
		response.HasError = true
		response.ErrorMessage = handlerResponse.ErrorMessage
		response.StatusCode = handlerResponse.StatusCode
	} else {
		response.HasError = false
		response.StatusCode = handlerResponse.StatusCode
		response.Id = uint32(handlerResponse.Data.(*model.Customer).ID)
		response.FirstName = handlerResponse.Data.(*model.Customer).FirstName
		response.SecondName = handlerResponse.Data.(*model.Customer).SecondName
		response.LastNameP = handlerResponse.Data.(*model.Customer).LastNameP
		response.LastNameM = handlerResponse.Data.(*model.Customer).LastNameM
		response.PhoneNumber = handlerResponse.Data.(*model.Customer).PhoneNumber
		response.Email = handlerResponse.Data.(*model.Customer).Email
	}
	return response
}
