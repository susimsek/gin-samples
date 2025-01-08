package dto

// LoginInput represents the login request input
// @Description Login request DTO containing username and password
type LoginInput struct {
	// Username is the username of the user
	Username string `json:"username" example:"admin" minLength:"3" maxLength:"50" validate:"required,min=3,max=50"`

	// Password is the password of the user
	Password string `json:"password" example:"password" minLength:"4" maxLength:"100" validate:"required,min=4,max=100"`
}
