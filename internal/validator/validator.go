package validator

import (
	"fmt"

	customerService "github.com/HouseCham/microservice-template/api/core/grpc"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// SetUpValidations sets up the custom validations for the validator
func SetUpValidations() {
	customValidator := validator.New()
	// Register custom validations to the validator
	customValidator.RegisterValidation("phone", validatePhone)
	customValidator.RegisterValidation("onlyAlpha", validateOnlyAlpha)
	customValidator.RegisterValidation("alpha_numeric", validateAlphaNumeric)
	customValidator.RegisterValidation("lettersAccents", validateLettersAndAccents)
	customValidator.RegisterValidation("lettersAccentsBlank", validateLettersAndAccentsBlank)

	Validate = customValidator
}

// GetValidatorErrorMessage returns a map with the validationError struct,
// specifying the field, tag and the error message
func GetValidatorErrorMessage(err error) []*customerService.ValidationError {
	var errors []*customerService.ValidationError
	if vErrs, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range vErrs {
			errors = append(errors, &customerService.ValidationError{
				Tag:   vErr.Tag(),
				Field: vErr.Field(),
				Err:   fmt.Sprintf(errorMessages[vErr.Tag()], validationFields[vErr.Field()]),
			})
		}
	}
	return errors
}