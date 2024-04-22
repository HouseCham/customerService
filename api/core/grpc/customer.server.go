package grpc

import (
	"context"

	"github.com/HouseCham/customerService/internal/handler"
	"github.com/HouseCham/customerService/internal/log"
)

type CustomerGrpcServer struct {
	CustomerServer
}

// InsertNewCostumer is a gRPC endpoint that receives a new costumer and inserts it into the database
func (s *CustomerGrpcServer) InsertCustomer(ctx context.Context, request *InsertCustomerCommand) (*InsertCustomerResponse, error) {
	log.Logger.Println("InsertCustomer gRPC endpoint invoked")

	// Convert gRPC request to model struct
	requestBody := mapInsertCustomerRequestToModel(request)

	// TODO: Add struct validation layer

	// Call the handler to insert the new customer
	handlerResponse := handler.InsertCustomerHandler(requestBody)

	// Convert handler response to gRPC response
	return mapHttpResponseToGrpc(handlerResponse, "InsertCustomerResponse").(*InsertCustomerResponse), nil
}
