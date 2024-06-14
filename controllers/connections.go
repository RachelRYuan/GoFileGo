package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"go-crud/utils"
)

func Connection() *gorm.DB {
	databaseURI := utils.GodotEnv("DATABASE_URL_DEV")

	db, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		logrus.Fatalf("Connection to Database Failed: %v", err)
		return nil
	}

	logrus.Info("Connection to Database Successfully")

	// Set up DB migrations here
	databaseMigrations(db)

	return db
}

func databaseMigrations(db *gorm.DB) {
	// Add your migration logic here

	logrus.Info("Database migrations completed")
}
