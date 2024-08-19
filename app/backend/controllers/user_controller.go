package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/JoshBenn/CTC-Coding-Challenge/database"
	"github.com/JoshBenn/CTC-Coding-Challenge/models"
	"github.com/JoshBenn/CTC-Coding-Challenge/services"
	"golang.org/x/crypto/bcrypt"
)

// For generically handling the user form structs
type user interface {
	// Gets a string of the user email
	GetEmail() string
	// Gets a string of the user password
	GetPassword() string
}

// Handles all registration and authentication requests
func AuthenticationHandler(node *common.Node) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(common.Authentication) {
			http.NotFound(writer, request)
			return
		}

		switch request.Method {
		// Registration
		case http.MethodPut:
			{
				// Parse the JSON body
				var registration models.RegisterRequest
				if err := json.NewDecoder(request.Body).Decode(&registration); err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				// Validate the data exists in the body
				var errs []string
				if len(registration.Username) == 0 {
					errs = append(errs, string(models.BlankUsername))
				}
				errs = append(errs, validateUser(&registration)...)

				// If any of the above failed then user cannot be created
				if len(errs) != 0 {
					http.Error(writer, strings.Join(errs, " "), http.StatusBadRequest)
					return
				}

				// Validate the email
				if err := validateEmail(&registration); err != nil {
					// Determine what type of error and return that
					if err.Error() == string(models.InvalidEmail) {
						http.Error(writer, string(models.InvalidEmail), http.StatusBadRequest)
						return
					}

					node.Log <- common.NewLog(common.Error, fmt.Sprintf("Unable to compile regex for email validation: %v", err))
					http.Error(writer, string(common.ServerCompilationError), http.StatusBadGateway)
					return
				}

				// Get the connection to the database
				provider, err := services.NewProvider(node)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadGateway)
					return
				}
				defer provider.CloseDbConn(node)

				// Check if the user already exists
				if err := queryUser(&provider, &registration); err != nil {
					http.Error(writer, models.UserError(models.UserExists), http.StatusBadRequest)
					return
				}

				// Hash the password
				hash, err := hashPassword(&registration)
				if err != nil {
					node.Log <- common.NewLog(common.Error, fmt.Sprintf("Could not hash password for %s\n\t%s", registration.Email, err.Error()))
					return
				}

				// Create the user
				createUserParams := database.CreateuserParams{
					Email:    registration.Email,
					Username: registration.Username,
					Password: hash,
				}
				user, err := provider.Queries.Createuser(context.Background(), createUserParams)
				if err != nil {
					node.Log <- common.NewLog(common.Error, fmt.Sprintf("Could not hash password for %s\n\t%s", registration.Email, err.Error()))
					return
				}

				// Return a response
				writer.Header().Set(string(common.ContentType), string(common.ApplicationJson))
				json.NewEncoder(writer).Encode(models.NewRegisterResponse(true, fmt.Sprintf("User with email <%s> created", user.Email)))

				node.Output <- fmt.Sprintf("User with email <%s> created", user.Email)
				node.Log <- common.NewLog(common.Info, fmt.Sprintf("User with email <%s> created", user.Email))

				break
			}

			// Log in/out
		case http.MethodPost:
			{
				// Parse the JSON body
				var login models.AuthenticationRequest
				if err := json.NewDecoder(request.Body).Decode(&login); err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				break
			}

			// All other methods passed
		default:
			{
				writer.Header().Set(string(common.Allow), fmt.Sprintf("%s, %s", common.Put, common.Post))
				http.Error(writer, string(common.MethodNotAllowed), http.StatusMethodNotAllowed)
			}
		}
	}
}

// Validates the user email and password
func validateUser[T user](user T) []string {
	var errs []string
	if len(user.GetEmail()) == 0 {
		errs = append(errs, string(models.BlankEmail))
	}
	if len(user.GetPassword()) == 0 {
		errs = append(errs, string(models.BlankPassword))
	}

	return errs
}

// Ensures the email is validated
func validateEmail[T user](user T) error {
	// Compile the regex
	reg, err := regexp.Compile(models.EMAIL_REGEX)
	if err != nil {
		return err
	}

	// Check the email to the regex
	if !reg.MatchString(user.GetEmail()) {
		return fmt.Errorf(strings.ToLower(string(models.InvalidEmail)))
	}

	// nil means valid
	return nil
}

func userToDb(user *models.User) database.User {
	return database.User{
		ID:       user.Id,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
}

// Converts a database user to a user struct
func dbToUser(dbUser *database.User) models.User {
	return models.User{
		Id:       dbUser.ID,
		Email:    dbUser.Email,
		Username: dbUser.Username,
		Password: dbUser.Password,
	}
}

// Gets all of the users currently existing in the database
func getUsers(provider *services.Provider) ([]models.User, error) {
	// Get all the database users
	dbUsers, err := provider.Queries.GetUsers(context.Background())
	if err != nil {
		return []models.User{}, err
	}

	// Convert to model users
	var users []models.User
	for _, dbUser := range dbUsers {
		users = append(users, dbToUser(&dbUser))
	}

	return users, nil
}

func getUserByEmail(provider *services.Provider, email string) (models.User, error) {
	// Get the user  via email
	dbUser, err := provider.Queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return models.User{}, err
	}

	return dbToUser(&dbUser), nil
}

func getUserByUsername(provider *services.Provider, username string) (models.User, error) {
	// Get the user via username
	dbUser, err := provider.Queries.GetUserByUsername(context.Background(), username)
	if err != nil {
		return models.User{}, err
	}

	return dbToUser(&dbUser), nil
}

// Query if the user exists via the email
func queryUser[T user](provider *services.Provider, user T) error {
	_, err := getUserByEmail(provider, user.GetEmail())
	return err
}

// Helpers from https://gowebexamples.com/password-hashing/
// Hashes the user's password using bcrypt at level 14
func hashPassword[T user](user T) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 14)
	return string(hash), err
}

// Compares the hash and a password
func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
