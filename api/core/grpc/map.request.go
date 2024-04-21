package grpc

import (
	"github.com/HouseCham/microservice-template/internal/model"
)

// Function to map InsertCustomerCommand to InsertCustomerModel
func mapInsertCustomerRequestToModel(request *InsertCustomerCommand) *model.InsertCustomerModel {
	return &model.InsertCustomerModel{
		FirstName:   request.FirstName,
		SecondName:  request.SecondName,
		LastNameP:   request.LastNameP,
		LastNameM:   request.LastNameM,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    request.Password,
	}
}