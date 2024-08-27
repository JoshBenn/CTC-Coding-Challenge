package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/JoshBenn/CTC-Coding-Challenge/models"
)

func ChatHandler(node *common.Node, chatroom *models.Chatroom) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(common.Authentication) {
			http.NotFound(writer, request)
			return
		}

		switch request.Method {
		// New message
		case http.MethodGet:
			{

				writer.Header().Set(string(common.ContentType), string(common.ApplicationJson))
				json.NewEncoder(writer).Encode(models.NewChatResponse(chatroom))
				break
			}

			// Get the messages
		case http.MethodPost:
			{

				break
			}

			// All other methods passed
		default:
			{
				writer.Header().Set(string(common.Allow), fmt.Sprintf("%s, %s", common.Put, common.Get))
				http.Error(writer, string(common.MethodNotAllowed), http.StatusMethodNotAllowed)
			}
		}
	}
}
