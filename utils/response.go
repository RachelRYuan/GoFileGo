package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Errors     interface{} `json:"errors"`
}

// APIResponse sends a JSON response with the given message, status code, method, and data.
func APIResponse(ctx *gin.Context, message string, statusCode int, method string, data interface{}) {
	jsonResponse := Response{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
		Data:       data,
	}

	ctx.JSON(statusCode, jsonResponse)
	if statusCode >= 400 {
		ctx.AbortWithStatus(statusCode)
	}
}

// ValidatorErrorResponse sends a JSON error response with the given status code, method, and error details.
func ValidatorErrorResponse(ctx *gin.Context, statusCode int, method string, err interface{}) {
	errResponse := ErrorResponse{
		StatusCode: statusCode,
		Method:     method,
		Errors:     err,
	}

	ctx.JSON(statusCode, errResponse)
	ctx.AbortWithStatus(statusCode)
}
