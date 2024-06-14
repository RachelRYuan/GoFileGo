package register

import (
	"time"
)

// RegisterResponse represents the structure of the response returned after a successful user registration.
type RegisterResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
