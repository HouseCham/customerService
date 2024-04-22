package grpc

import (
	"github.com/HouseCham/customerService/internal/model"
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

// mapHttpResponseToGrpc is a function to map the response from the handler to the gRPC response
func mapHttpResponseToGrpc(response model.HttpResponse, returnStruct string) interface{} {
	switch returnStruct {
	case "InsertCustomerResponse":
		newId := uint32(0)
		if response.Data != nil {
			newId = uint32(response.Data.(uint))
		}
		return &InsertCustomerResponse{
			StatusCode: response.StatusCode,
			NewId:        newId,
			HasError:     response.HasError,
			ErrorMessage: response.ErrorMessage,
		}
	default:
		// Handle other response types here, if needed
		return nil
	}
}
