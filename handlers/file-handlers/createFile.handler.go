package filehandlers

import (
	filecontrollers "GOFILEGO/controllers/file-controllers"
	"GOFILEGO/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateHandler handles the creation and upload of a file.
func (h *handler) CreateHandler(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		utils.APIResponse(context, "Failed to get file from request", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	defer file.Close()

	result, err := utils.UploadFile(file, header.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println(err)
		utils.APIResponse(context, "Failed to upload file", http.StatusInternalServerError, http.MethodPost, nil)
		return
	}

	fileInput := filecontrollers.FileInput{
		ID:   result.PublicID,
		Type: result.Format,
		Name: header.Filename,
	}

	fileResponse, statusCode := h.service.CreateFile(&fileInput)

	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Uploaded the file successfully.", http.StatusCreated, http.MethodPost, fileResponse)
	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occurred", http.StatusExpectationFailed, http.MethodPost, nil)
	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusConflict, http.MethodPost, nil)
	default:
		utils.APIResponse(context, "Unexpected error", http.StatusInternalServerError, http.MethodPost, nil)
	}
}
