package handler

import (
	"github.com/HouseCham/customerService/internal/log"
	"github.com/HouseCham/customerService/internal/model"
	"google.golang.org/grpc/codes"
)

// handleErrorResponse is a function that logs an error and returns an error response
func handleErrorResponse(message string, err error, statusCode codes.Code) model.HttpResponse {
    log.Logger.Errorf("%s: %v", message, err)
    return model.HttpResponse{
        StatusCode:   uint32(statusCode),
        Data:         nil,
        HasError:     true,
        ErrorMessage: err.Error(),
    }
}