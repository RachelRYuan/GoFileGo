package filehandlers

import (
	filecontrollers "GOFILEGO/controllers/file-controllers"
	"GOFILEGO/models"
	"GOFILEGO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateHandler handles file upload requests.
func (h *handler) CreateHandler(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		utils.APIResponse(context, "Invalid file upload request", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	defer file.Close()

	var user models.UserEntity
	jwtData, exists := context.Get("user")
	if !exists {
		utils.APIResponse(context, "Unauthorized user", http.StatusUnauthorized, http.MethodPost, nil)
		return
	}

	// Convert JWT data to user entity
	err = utils.StringToEntity(jwtData, &user)
	if err != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, http.MethodPost, nil)
		return
	}

	// Upload the file
	result, err := utils.UploadFile(file, header.Header.Get("Content-Type"))
	if err != nil {
		utils.APIResponse(context, "Unable to upload file to the server", http.StatusFailedDependency, http.MethodPost, nil)
		return
	}

	// Prepare file input
	fileInput := filecontrollers.FileInput{
		ID:     result.PublicID,
		Type:   result.Format,
		Name:   header.Filename,
		UserId: user.ID,
	}

	// Create file entry in the database
	fileResponse, statusCode := h.service.CreateFile(&fileInput)

	// Handle response based on the status code
	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Uploaded the file successfully.", http.StatusCreated, http.MethodPost, fileResponse)

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occurred", http.StatusExpectationFailed, http.MethodPost, nil)

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusConflict, http.MethodPost, nil)

	default:
		utils.APIResponse(context, "An unknown error occurred", http.StatusInternalServerError, http.MethodPost, nil)
	}
}
