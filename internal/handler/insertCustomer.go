package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"github.com/HouseCham/customerService/internal/repository"
	"github.com/HouseCham/customerService/internal/validator"
	util "github.com/HouseCham/customerService/pkg/utils"
	"google.golang.org/grpc/codes"
)

// InsertCustomerHandler is a handler that receives a new customer and inserts it into the database
func InsertCustomerHandler(requestBody *model.Customer) model.HttpResponse {
    log.Logger.Infoln("InsertCustomerHandler invoked")

    // Create a channel to communicate the result back
    resultCh := make(chan model.HttpResponse)

    // Perform the operations in a goroutine
    go func() {
        defer close(resultCh)

        // Validate the request body
        if err := validator.Validate.Struct(requestBody); err != nil {
            resultCh <- handleErrorResponse("Error validating request body", err, codes.InvalidArgument, validator.GetValidatorErrorMessage(err))
            return
        }

        // Hash the password
        hashedPassword, err := util.HashPassword(requestBody.Password)
        if err != nil {
            resultCh <- handleErrorResponse("Error hashing password", err, codes.Internal, nil)
            return
        }
        requestBody.HashedPassword = hashedPassword

        // Call the repository to insert the new customer
        newId, err := repository.InsertNewCustomerDB(requestBody)
        if err != nil {
            resultCh <- handleErrorResponse("Error inserting customer", err, codes.Internal, nil)
            return
        }

        // Send the successful response
        resultCh <- model.HttpResponse{
            StatusCode:   uint32(codes.OK),
            Data:         newId,
            HasError:     false,
            ErrorMessage: "",
        }
    }()

    // Wait for the result from the goroutine
    return <-resultCh
}
