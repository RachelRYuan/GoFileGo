package filecontrollers

import (
	"GOFILEGO/models"
)

// Service defines the methods that any file service must implement.
type Service interface {
	CreateFile(input *FileInput) (*models.FileModel, int)
	GetAllFiles(userId uint) ([]models.FileModel, int)
	DeleteFile(fileID uint) int
}

// service struct implements the Service interface.
type service struct {
	repository Repository
}

// NewFileService initializes a new file service with the provided repository.
func NewFileService(r Repository) *service {
	return &service{repository: r}
}

// CreateFile creates a new file record in the database.
func (s *service) CreateFile(input *FileInput) (*models.FileModel, int) {
	fileModel := models.FileModel{
		Type:      input.Type,
		Name:      input.Name,
		Url:       input.Url,
		AccessKey: input.ID,
		UserID:    input.UserId,
	}
	return s.repository.CreateFile(&fileModel)
}

// GetAllFiles retrieves all files associated with a specific user.
func (s *service) GetAllFiles(userId uint) ([]models.FileModel, int) {
	return s.repository.GetAllFiles(userId)
}

// DeleteFile deletes a file record from the database.
func (s *service) DeleteFile(fileID uint) int {
	return s.repository.DeleteFile(fileID)
}
