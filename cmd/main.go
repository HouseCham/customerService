package main

import (
	"fmt"
	"net"

	grpcService "github.com/HouseCham/customerService/api/core/grpc"
	"github.com/HouseCham/customerService/internal/config"
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/repository"
	"github.com/HouseCham/customerService/internal/validator"
	"google.golang.org/grpc"
)

func main() {
	// Setting up logger
	log.SetUpLogger()

	// Setting up config file
	log.Logger.Println("Setting up config file")
	err := config.GetConfigFile()
	if err != nil {
		log.Logger.Panicf("Failed to get config file: %v", err)
	}

	// Setting up database connection
	log.Logger.Println("Setting up database connection")
	err = repository.SetUpDBConnection()
	if err != nil {
		log.Logger.Panicf("Failed to set up database connection: %v", err)
	}

	// Create a listener on the configured port
	log.Logger.Printf("Starting gRPC listener on port %d", config.ConfigFile.App.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ConfigFile.App.Port))
	if err != nil {
		log.Logger.Errorf("Failed to listen: %v", err)
		panic(err)
	}

	// Create a new gRPC server
	log.Logger.Println("Creating a new gRPC server")
	grpcServer := grpc.NewServer()

	// Register your gRPC service implementation with the server
	log.Logger.Println("Starting to implement gRPC service")
	grpcService.RegisterCustomerServer(grpcServer, &grpcService.CustomerGrpcServer{})

	// Setting up custom validations
	log.Logger.Println("Setting up custom validations")
	validator.SetUpValidations()

	// Start the gRPC server
	log.Logger.Printf("gRPC server is running on port %d", config.ConfigFile.App.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Logger.Fatalf("Failed to serve: %v", err)
	}
}
