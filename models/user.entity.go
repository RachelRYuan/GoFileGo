package models

import (
	"time"
	"GOFILEGO/utils"

	"github.com/jinzhu/gorm"
)

type UserEntity struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"column:username;unique;not null" json:"username"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Image     *string   `gorm:"column:image" json:"image,omitempty"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate is a GORM hook that is called before a new record is inserted into the database.
func (entity *UserEntity) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(entity.Password)
	if err != nil {
		return err
	}
	entity.Password = hashedPassword
	entity.CreatedAt = time.Now().Local()
	return nil
}

// BeforeUpdate is a GORM hook that is called before an existing record is updated in the database.
func (entity *UserEntity) BeforeUpdate(tx *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
