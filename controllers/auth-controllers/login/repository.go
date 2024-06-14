package loginAuth

import (
	"GOFILEGO/models"
	"GOFILEGO/utils"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Repository interface defines a contract for the repository responsible for handling login-related operations.
// It has one function `LoginRepository` which takes the UserEntity and returns the UserEntity and an int status code.
type Repository interface {
	LoginRepository(input *models.UserEntity) (*models.UserEntity, int)
}

// repository struct is the concrete implementation of the Repository interface.
type repository struct {
	db *gorm.DB
}

// NewRepositoryLogin function is the constructor for the repository struct that creates a new instance of the repository struct.
func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

// LoginRepository method of the repository struct implements the LoginRepository method from the Repository interface.
func (r *repository) LoginRepository(input *models.UserEntity) (*models.UserEntity, int) {
	// Check if the user exists
	var user models.UserEntity
	db := r.db.Model(&user)
	checkAccount := db.Select("*").Where("email = ?", input.Email).First(&user)

	if checkAccount.Error != nil {
		if checkAccount.Error == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound
		}
		logrus.Error("Database error: ", checkAccount.Error)
		return nil, http.StatusInternalServerError
	}

	// Check if the password matches
	if err := utils.ComparePassword(user.Password, input.Password); err != nil {
		return nil, http.StatusUnauthorized
	}

	return &user, http.StatusAccepted
}
