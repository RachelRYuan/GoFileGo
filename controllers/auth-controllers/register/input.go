package register

// RegisterInput represents the expected structure for register request payloads,
// including validation rules for each field.
type RegisterInput struct {
	Username string `json:"username" validate:"required,lowercase"` // The user's username, required and must be in lowercase.
	Email    string `json:"email" validate:"required,email"`        // The user's email address, required and must be a valid email format.
	Password string `json:"password" validate:"required,gte=8"`     // The user's password, required and must be at least 8 characters long.
}
