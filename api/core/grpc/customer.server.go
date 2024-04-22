package grpc

import (
	"context"
	"log"

	"github.com/HouseCham/microservice-template/internal/handler"
	"github.com/fatih/color"
	codes "google.golang.org/grpc/codes"
)

type CustomerGrpcServer struct {
	CustomerServer
}

// InsertNewCostumer is a gRPC endpoint that receives a new costumer and inserts it into the database
func (s *CustomerGrpcServer) InsertCustomer(ctx context.Context, request *InsertCustomerCommand) (*InsertCustomerResponse, error) {
	log.Println(color.BlueString("InsertCustomer gRPC endpoint invoked"))

	// Convert gRPC request to model struct
	requestBody := mapInsertCustomerRequestToModel(request)

	// TODO: Add struct validation layer

	// Call the handler to insert the new customer
	newId, statusCode, err := handler.InsertCustomerHandler(requestBody)

	if err != nil {
		log.Printf(color.RedString("Error inserting customer: %v"), err)
		return &InsertCustomerResponse{
			HasError:     true,
			ErrorMessage: "Failed to insert customer",
			NewId:        0,
			StatusCode:   uint32(codes.Internal),
		}, nil
	}

	return &InsertCustomerResponse{
		HasError:     false,
		ErrorMessage: "",
		NewId:        uint32(newId),
		StatusCode:   uint32(statusCode),
	}, nil
}