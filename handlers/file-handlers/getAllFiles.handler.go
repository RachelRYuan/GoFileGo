package fileHandlers

import (
	"GOFILEGO/models"
	"GOFILEGO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllFilesHandler handles requests to retrieve all files for a user.
func (h *handler) GetAllFilesHandler(context *gin.Context) {
	jwtData, exists := context.Get("user")
	if !exists {
		utils.APIResponse(context, "Unauthorized user", http.StatusUnauthorized, http.MethodPost, nil)
		return
	}

	var user models.UserEntity
	// Convert JWT data to user entity
	err := utils.StringToEntity(jwtData, &user)
	if err != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, http.MethodPost, nil)
		return
	}

	// Retrieve all files for the user
	fileResponse, statusCode := h.service.GetAllFiles(user.ID)

	// Handle response based on the status code
	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Received files", http.StatusOK, http.MethodPost, fileResponse)

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal server error occurred", http.StatusExpectationFailed, http.MethodPost, nil)

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusConflict, http.MethodPost, nil)

	default:
		utils.APIResponse(context, "An unknown error occurred", http.StatusInternalServerError, http.MethodPost, nil)
	}
}
