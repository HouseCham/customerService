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
	log.Logger.Info("Setting up config file")
	if err := config.GetConfigFile(); err != nil {
		log.Logger.Fatalf("Failed to get config file: %v", err)
	}

	// Setting up database connection
	log.Logger.Info("Setting up database connection")
	if err := repository.SetUpDBConnection(); err != nil {
		log.Logger.Fatalf("Failed to set up database connection: %v", err)
	}

	// Create a listener on the configured port
	log.Logger.Infof("Starting gRPC listener on port %d", config.ConfigFile.App.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ConfigFile.App.Port))
	if err != nil {
		log.Logger.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	log.Logger.Info("Creating a new gRPC server")
	grpcServer := grpc.NewServer()

	// Register your gRPC service implementation with the server
	log.Logger.Info("Registering gRPC service")
	grpcService.RegisterCustomerServer(grpcServer, &grpcService.CustomerGrpcServer{})

	// Setting up custom validations
	log.Logger.Info("Setting up custom validations")
	validator.SetUpValidations()

	// Start the gRPC server
	log.Logger.Infof("gRPC server is running on port %d", config.ConfigFile.App.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Logger.Fatalf("Failed to serve: %v", err)
	}
}