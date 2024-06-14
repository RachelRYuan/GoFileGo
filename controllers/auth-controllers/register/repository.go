package register

import (
	"GOFILEGO/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Repository interface defines a contract for the repository responsible for handling registration-related operations.
type Repository interface {
	RegisterRepository(input *models.UserEntity) (*models.UserEntity, int)
}

// repository struct is the concrete implementation of the Repository interface.
type repository struct {
	db *gorm.DB
}

// NewRegisterRepository function is the constructor for the repository struct that creates a new instance of the repository struct.
func NewRegisterRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// RegisterRepository method of the repository struct implements the RegisterRepository method from the Repository interface.
func (r *repository) RegisterRepository(input *models.UserEntity) (*models.UserEntity, int) {
	var user models.UserEntity
	db := r.db

	// Check if user exists
	checkUserAccount := db.Select("*").Where("email = ?", input.Email).Find(&user)
	if checkUserAccount.RowsAffected > 0 {
		return nil, http.StatusConflict
	}

	// If not, create the user in the database
	if err := db.Create(&input).Error; err != nil {
		logrus.Error("Failed to create user: ", err)
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}
