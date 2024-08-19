package controllers

import (
	"fmt"
	"net/http"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
)

func ChatHandler(node *common.Node) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(common.Authentication) {
			http.NotFound(writer, request)
			return
		}

		switch request.Method {
		// Registration
		case http.MethodPut:
			{
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
