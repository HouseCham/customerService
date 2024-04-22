package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"github.com/HouseCham/customerService/internal/repository"
	util "github.com/HouseCham/customerService/pkg/utils"
	"google.golang.org/grpc/codes"
)

// InsertCustomerHandler is a handler that receives a new customer and inserts it into the database
func InsertCustomerHandler(requestBody *model.Customer) model.HttpResponse {
	log.Logger.Infoln("InsertCustomerHandler invoked")
    // Hash the password
    hashedPassword, err := util.HashPassword(requestBody.Password)
    if err != nil {
        return handleErrorResponse("Error hashing password", err, codes.Internal)
    }
    requestBody.HashedPassword = hashedPassword

    // Call the repository to insert the new customer
    newId, err := repository.InsertNewCostumerDB(requestBody)
    if err != nil {
        return handleErrorResponse("Error inserting customer", err, codes.Internal)
    }

    return model.HttpResponse{
        StatusCode:   uint32(codes.OK),
        Data:         newId,
        HasError:     false,
        ErrorMessage: "",
    }
}