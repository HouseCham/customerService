package grpc

import (
	"context"

	"github.com/HouseCham/customerService/internal/handler"
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
)

type CustomerGrpcServer struct {
	CustomerServer
}

// InsertNewCostumer is a gRPC endpoint that receives a new costumer and inserts it into the database
func (s *CustomerGrpcServer) InsertCustomer(ctx context.Context, request *InsertCustomerCommand) (*InsertCustomerResponse, error) {
	log.Logger.Println("InsertCustomer gRPC endpoint invoked")

	// Convert gRPC request to model struct
	requestBody := mapInsertCustomerRequestToModel(request)

	// Call the handler to insert the new customer
	handlerResponse := handler.InsertCustomerHandler(requestBody)

	// Check if there was an error
	response := handleInsertResponseErrors(handlerResponse)

	log.Logger.Println("InsertCustomer gRPC endpoint finished")

	return response, nil
}

// UpdateCustomer is a gRPC endpoint that receives a customer and updates it in the database
func (s *CustomerGrpcServer) UpdateCustomer(ctx context.Context, request *UpdateCustomerCommand) (*UpdateCustomerResponse, error) {
	log.Logger.Println("UpdateCustomer gRPC endpoint invoked")

	// Convert gRPC request to model struct
	requestBody := mapUpdateCustomerRequestToModel(request)

	// Call the handler to update the customer
	handlerResponse := handler.UpdateCustomerHandler(requestBody)

	// Check if there was an error
	response := handleUpdateResponseErrors(handlerResponse)

	log.Logger.Println("UpdateCustomer gRPC endpoint finished")

	return response, nil
}

// DeleteCustomer is a gRPC endpoint that receives a customer ID and deletes it from the database
func (s *CustomerGrpcServer) DeleteCustomer(ctx context.Context, request *DeleteCustomerCommand) (*DeleteCustomerResponse, error) {
	log.Logger.Println("DeleteCustomer gRPC endpoint invoked")

	// Call the handler to delete the customer
	handlerResponse := handler.DeleteCustomerHandler(uint(request.Id))

	log.Logger.Println("DeleteCustomer gRPC endpoint finished")
	return &DeleteCustomerResponse{
		HasError:     handlerResponse.HasError,
		ErrorMessage: handlerResponse.ErrorMessage,
		StatusCode:   handlerResponse.StatusCode,
	}, nil
}

// GetCustomer is a gRPC endpoint that receives a customer ID and returns the customer from the database
func (s *CustomerGrpcServer) GetCustomer(ctx context.Context, request *GetCustomerQuery) (*GetCustomerResponse, error) {
	log.Logger.Println("GetCustomer gRPC endpoint invoked")

	// Call the handler to get the customer
	handlerResponse := handler.GetCustomerByIdHandler(uint(request.Id))

	log.Logger.Println("GetCustomer gRPC endpoint finished")

	return &GetCustomerResponse{
		HasError:     handlerResponse.HasError,
		ErrorMessage: handlerResponse.ErrorMessage,
		StatusCode:   handlerResponse.StatusCode,

		Id: 		uint32(handlerResponse.Data.(*model.Customer).ID),
		FirstName: 	handlerResponse.Data.(*model.Customer).FirstName,
		SecondName: handlerResponse.Data.(*model.Customer).SecondName,
		LastNameP: 	handlerResponse.Data.(*model.Customer).LastNameP,
		LastNameM: 	handlerResponse.Data.(*model.Customer).LastNameM,
		PhoneNumber: handlerResponse.Data.(*model.Customer).PhoneNumber,
		Email: 		handlerResponse.Data.(*model.Customer).Email,
	}, nil
}