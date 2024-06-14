package utils

import (
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Results struct {
		Errors []map[string]map[string]string `json:"errors"`
	} `json:"results"`
}

type ErrorMetaConfig struct {
	Tag     string
	Field   string
	Message string
}

// GoValidator validates the given struct using the provided configuration.
func GoValidator(s interface{}, config []ErrorMetaConfig) (interface{}, int) {
	validate := validator.New()
	err := validate.Struct(s)
	if err == nil {
		return nil, 0
	}

	validationErrors := err.(validator.ValidationErrors)
	errorMessages := []map[string]map[string]string{}

	for _, validationError := range validationErrors {
		fieldError := map[string]map[string]string{
			validationError.Field(): {
				validationError.Tag(): getErrorMessage(validationError, config),
			},
		}
		errorMessages = append(errorMessages, fieldError)
	}

	return errorMessages, len(validationErrors)
}

// getErrorMessage returns the custom error message if defined in the config.
func getErrorMessage(err validator.FieldError, config []ErrorMetaConfig) string {
	for _, cfg := range config {
		if cfg.Field == err.Field() && cfg.Tag == err.Tag() {
			return cfg.Message
		}
	}
	return err.Error()
}
