package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/JoshBenn/CTC-Coding-Challenge/models"
)

type user interface {
	GetEmail() string
	GetPassword() string
}

type register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r register) GetEmail() string {
	return r.Email
}
func (r register) GetPassword() string {
	return r.Email
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l login) GetEmail() string {
	return l.Email
}
func (l login) GetPassword() string {
	return l.Email
}

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
				var registration register
				if err := json.NewDecoder(request.Body).Decode(&registration); err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				// Validate the data exists in the body
				var errs []string
				if len(registration.Email) == 0 {
					errs = append(errs, string(models.BlankEmail))
				}
				if len(registration.Username) == 0 {
					errs = append(errs, string(models.BlankUsername))
				}
				if len(registration.Password) == 0 {
					errs = append(errs, string(models.BlankPassword))
				}

				// If any of the above failed then user cannot be created
				if len(errs) != 0 {
					http.Error(writer, strings.Join(errs, " "), http.StatusBadRequest)
					return
				}

				// Validate the email
				if err := validateEmail(registration); err != nil {
					// Determine what type of error and return that
					if err.Error() == string(models.InvalidEmail) {
						http.Error(writer, string(models.InvalidEmail), http.StatusBadRequest)
						return
					}

					node.Log <- common.NewLog(common.Error, fmt.Sprintf("Unable to compile regex for email validation: %v", err))
					http.Error(writer, string(common.ServerCompilationError), http.StatusBadGateway)
					return
				}

				break
			}

			// Log in/out
		case http.MethodPost:
			{
				// Parse the JSON body
				var login login
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

// Ensures the email is validated
func validateEmail[T user](user T) error {
	// Compile the regex
	reg, err := regexp.Compile(models.EMAIL_REGEX)
	if err != nil {
		return err
	}

	// Check the email to the regex
	if !reg.MatchString(user.GetEmail()) {
		return fmt.Errorf(string(models.InvalidEmail))
	}

	// nil means valid
	return nil
}
