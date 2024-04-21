package main

import (
	"fmt"
	"log"
	"net"

	"github.com/HouseCham/microservice-template/internal/config"
	"github.com/HouseCham/microservice-template/internal/repository"
	"github.com/HouseCham/microservice-template/internal/validator"
    customerService "github.com/HouseCham/microservice-template/api/core/grpc"
	"google.golang.org/grpc"
)

func main() {
    // Setting up config file
    err := config.GetConfigFile()
    if err != nil {
        panic(err)
    }

    // Setting up database connection
    err = repository.SetUpDBConnection()
    if err != nil {
        panic(err)
    }

    // Create a listener on the configured port
    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ConfigFile.App.Port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register your gRPC service implementation with the server
    customerService.RegisterCustomerServer(grpcServer, &customerService.CustomerGrpcServer{})

    // Setting up custom validations
    validator.SetUpValidations()

    // Start the gRPC server
    log.Printf("gRPC server is running on port %d", config.ConfigFile.App.Port)
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
