package filehandlers

import (
	filecontrollers "GOFILEGO/controllers/file-controllers"
)

// handler struct for managing file-related operations
type handler struct {
	service filecontrollers.Service
}

// NewCreateHandler initializes a new handler with the provided file service
func NewCreateHandler(service filecontrollers.Service) *handler {
	return &handler{service: service}
}
