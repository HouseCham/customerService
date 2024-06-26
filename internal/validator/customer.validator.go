package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// validatePhone validates that the phone number inserted is only numbers
func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if phone == "" {
		return true
	}
	match, _ := regexp.MatchString("^[0-9]+$", phone)
	return match
}

// validateOnlyAlpha validates that the field inserted is only letters
func validateOnlyAlpha(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	if name == "" {
		return true
	}
	match, _ := regexp.MatchString("^[a-zA-Z]+$", name)
	return match
}

// validateLettersAndAccents validates that the field inserted is only letters and accents
func validateLettersAndAccents(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}
	regex := regexp.MustCompile("^[a-zA-ZÀ-ÿ]+$") // Regular expression to match letters and accents
	return regex.MatchString(value)
}

// validateAlphaNumeric validates that the field inserted is only letters and numbers
func validateAlphaNumeric(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	if name == "" {
		return true
	}
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", name)
	return match
}

// validateLettersAndAccentsBlank validates that the field inserted is only letters and accents
// and can there can be blank spaces also
func validateLettersAndAccentsBlank(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	regex := regexp.MustCompile(`^[A-Za-zÀ-ÿ ]+$`)
	return regex.MatchString(input)
}