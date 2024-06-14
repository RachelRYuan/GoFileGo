package configs

import (
	"GOFILEGO/models"
	"GOFILEGO/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

// Connection establishes a connection to the database using the provided environment variable.
func Connection() *gorm.DB {
	databaseURI := utils.GodotEnv("DATABASE_URL_DEV")

	db, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		logrus.Fatalf("Connection to Database Failed: %v", err)
	} else {
		logrus.Info("Connection to Database Successfully")
	}

	// Setup database migrations
	databaseMigrations(db)

	return db
}

// databaseMigrations performs the database migrations.
func databaseMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.UserEntity{})
	logrus.Info("Database migrations completed")
}
