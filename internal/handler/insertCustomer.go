package handler

import (
	"github.com/HouseCham/microservice-template/internal/model"
)

// InsertCustomerHandler is a handler that receives a new customer and inserts it into the database
func InsertCustomerHandler(requestBody *model.InsertCustomerModel) (int, int, error) {
	// Call the repository to insert the new customer
	// newId, statusCode, err := repository.InsertCustomerRepository(requestBody)
	// if err != nil {
	// 	return 0, statusCode, err
	// }
	// return newId, statusCode, nil
	return 0, 0 , nil
}