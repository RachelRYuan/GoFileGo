package filehandlers

import (
	filecontrollers "GOFILEGO/controllers/file-controllers"
	"encoding/json"
	"net/http"
)

// Handler struct contains the service to be used by the handler functions.
type handler struct {
	service filecontrollers.Service
}

// NewHandler initializes a new handler with the given service.
func NewHandler(service filecontrollers.Service) *handler {
	return &handler{service: service}
}

// CreateFileHandler handles the creation of a new file.
func (h *handler) CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	var input filecontrollers.FileInput

	// Decode the request body into the FileInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service to create a new file
	file, status := h.service.CreateFile(&input)
	if status != http.StatusCreated {
		http.Error(w, "Failed to create file", status)
		return
	}

	// Return the created file as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(file); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
