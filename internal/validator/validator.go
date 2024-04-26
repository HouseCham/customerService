package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/HouseCham/customerService/internal/model"
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
func GetValidatorErrorMessage(err error) []model.ValidatorError {
	var errors []model.ValidatorError
	if vErrs, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range vErrs {
			errors = append(errors, model.ValidatorError{
				Tag:   vErr.Tag(),
				Field: vErr.Field(),
				Err:   fmt.Sprintf(errorMessages[vErr.Tag()], validationFields[vErr.Field()]),
			})
		}
	}
	return errors
}