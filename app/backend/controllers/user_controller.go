package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
)

const EMAIL_REGEX = `/^[^\s@]+@[^\s@]+\.[^\s@]+$/`

type register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

				// Validate the data exists

				break
			}

			// Log in/out
		case http.MethodPost:
			{

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
