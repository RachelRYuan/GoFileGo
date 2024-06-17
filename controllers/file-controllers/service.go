package filecontrollers

import (
	"GOFILEGO/models"
	"net/http"
)

// FileInput represents the input data for creating a file.
type FileInput struct {
	Type string
	Name string
	Url  string
}

// Service interface defines the methods that the file service should implement.
type Service interface {
	CreateFile(input *FileInput) (*models.FileModel, int)
	GetAllFiles() ([]models.FileModel, int)
	DeleteFile(fileID uint) int
}

// service struct implements the Service interface.
type service struct {
	repository Repository
}

// NewFileService creates a new instance of the service.
func NewFileService(r Repository) Service {
	return &service{repository: r}
}

// CreateFile processes the input and calls the repository to create a new file record.
func (s *service) CreateFile(input *FileInput) (*models.FileModel, int) {
	// Create a new FileModel from the input data
	fileModel := models.FileModel{
		Type: input.Type,
		Name: input.Name,
		Url:  input.Url,
	}

	// Call the repository to create the file
	createdFile, status := s.repository.CreateFile(&fileModel)
	if status != http.StatusCreated {
		return nil, status
	}

	return createdFile, http.StatusCreated
}

// GetAllFiles retrieves all file records by calling the repository.
func (s *service) GetAllFiles() ([]models.FileModel, int) {
	files, status := s.repository.GetAllFiles()
	if status != http.StatusOK {
		return nil, status
	}

	return files, http.StatusOK
}

// DeleteFile deletes a file record by ID by calling the repository.
func (s *service) DeleteFile(fileID uint) int {
	status := s.repository.DeleteFile(fileID)
	if status != http.StatusNoContent {
		return status
	}

	return http.StatusNoContent
}
