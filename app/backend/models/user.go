package models

// Models version of the user struct for json-ification
type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Generates a new user
func NewUser(id int64, email, username, password string) User {
	return User{
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
	}
}

// For easy error handling and message passing, this could be more robust with structs
type UserErr string

const (
	InvalidEmail       UserErr = "Invalid Email"
	BlankEmail         UserErr = "Email cannot be blank"
	InvalidUsername    UserErr = "Invalid Username"
	BlankUsername      UserErr = "Username cannot be blank"
	InvalidPassword    UserErr = "Invalid Password"
	BlankPassword      UserErr = "Password cannot be blank"
	DoesNotExist       UserErr = "User does not exist"
	InvalidCredentials UserErr = "InvalidCredentials"
	UserExists         UserErr = "User already exists"

	EMAIL_REGEX = `/^[^\s@]+@[^\s@]+\.[^\s@]+$/`
)

// Create the error with this string
func UserError(err UserErr) string {
	return string(err)
}

// Registration form
type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RegisterRequest) GetEmail() string {
	return r.Email
}
func (r *RegisterRequest) GetPassword() string {
	return r.Email
}

// Primary struct for all responses
type registerResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewRegisterResponse(ok bool, message string) registerResponse {
	return registerResponse{
		Ok: ok, Message: message,
	}
}

// Authentication form
type AuthenticationRequest struct {
	InOut    string `json:"in-out"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *AuthenticationRequest) GetEmail() string {
	return l.Email
}
func (l *AuthenticationRequest) GetPassword() string {
	return l.Email
}

type authenticationResponse struct {
}

func NewAuthenticationResponse() authenticationResponse {
	return authenticationResponse{}
}
