package filecontrollers

import (
	"GOFILEGO/models"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Repository interface defines the methods that any type that wants to be a repository must implement.
type Repository interface {
	CreateFile(file *models.FileModel) (*models.FileModel, int)
	GetAllFiles() ([]models.FileModel, int)
	DeleteFile(fileID uint) int
}

// repository struct implements the Repository interface.
type repository struct {
	db *gorm.DB
}

// NewFileRepository creates a new instance of repository.
func NewFileRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// CreateFile inserts a new file record into the database.
func (repo *repository) CreateFile(file *models.FileModel) (*models.FileModel, int) {
	if err := repo.db.Create(file).Error; err != nil {
		return nil, http.StatusInternalServerError
	}
	return file, http.StatusCreated
}

// GetAllFiles retrieves all file records from the database.
func (repo *repository) GetAllFiles() ([]models.FileModel, int) {
	var files []models.FileModel
	if err := repo.db.Find(&files).Error; err != nil {
		return nil, http.StatusInternalServerError
	}
	return files, http.StatusOK
}

// DeleteFile deletes a file record from the database by ID.
func (repo *repository) DeleteFile(fileID uint) int {
	if err := repo.db.Where("id = ?", fileID).Delete(&models.FileModel{}).Error; err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusNoContent
}
