package models

import "time"

// FileModel represents the schema for the files table.
type FileModel struct {
	ID        uint      `gorm:"primary_key"`    // Unique identifier for the file
	Type      string    `gorm:"not null"`       // Type of the file (e.g., image, video)
	Name      string    `gorm:"not null"`       // Name of the file
	Url       string    `gorm:"not null"`       // URL where the file is stored
	AccessKey string    `gorm:""`               // Optional access key for the file
	CreatedAt time.Time `gorm:"autoCreateTime"` // Timestamp when the file was created
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Timestamp when the file was last updated
}
