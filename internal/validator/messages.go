package validator

// errorMessages is the map that contains the error messages
// that will be returned when a validation fails
var errorMessages = map[string]string{
	"required":            "The field %s is required.",
	"phone":               "The field %s must be a valid phone number.",
	"onlyAlpha":           "The field %s can only contain letters.",
	"lettersAccents":      "The field %s can only contain letters and accents.",
	"lettersAccentsBlank": "The field %s can only contain letters, accents and blank spaces.",
	"email":               "The field %s must be a valid email address.",
	"max":                 "The field %s exceeds the maximum length of characters allowed.",
	"min":                 "The field %s is below the minimum length of characters allowed.",
	"numeric":             "The field %s must be a valid number.",
}

// validationFields is the map that contains the fields that will be validated
var validationFields = map[string]string{
	"FirstName": "name",
	"LastNameP": "father's last name",
	"LastNameM": "mother's last name",
	"Email":     "email",
	"Phone":     "phone",
	"Password":  "password",

	/* ===== Address ===== */
	"Street":    "street",
	"Suburb":    "suburb",
	"State":     "state",
	"City":      "city",
	"ZipCode":   "zip code",
	"ExtNumber": "exterior number",
	"IntNumber": "interior number",
	"Reference": "reference",
	"Country":   "country",
}