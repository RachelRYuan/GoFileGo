package loginAuth

// LoginInput represents the expected structure for login request payloads,
// including validation rules for each field.
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"` // The user's email address, required and must be a valid email format.
	Password string `json:"password" validate:"required"`    // The user's password, required.
}
