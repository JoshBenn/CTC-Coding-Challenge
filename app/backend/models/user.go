package models

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
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
)

// Create the error with this string
func UserError(err UserErr) string {
	return string(err)
}
