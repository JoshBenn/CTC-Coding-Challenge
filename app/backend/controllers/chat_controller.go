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
		if request.URL.Path != string(common.Chat) {
			http.NotFound(writer, request)
			return
		}

		switch request.Method {
		// New message
		case http.MethodGet:
			{
				writer.Header().Set(string(common.ContentType), string(common.ApplicationJson))
				if err := json.NewEncoder(writer).Encode(models.NewChatResponse(chatroom)); err != nil {
					node.Log <- common.NewLog(common.Error, err.Error())
				}
				return
			}

			// Get the messages
		case http.MethodPost:
			{
				fmt.Println(request.Body)
				var msg models.Message
				decoder := json.NewDecoder(request.Body)
				if err := decoder.Decode(&msg); err != nil {
					node.Output <- fmt.Sprintf("Could not parse body %v", request.Body)
					http.Error(writer, "Invalid message format", http.StatusBadRequest)
					return
				}

				// Message validation
				if msg.Username == "" || msg.Content == "" {
					node.Output <- "Username and content cannot be empty"
					http.Error(writer, "Username and content cannot be empty", http.StatusBadRequest)
					return
				}

				// Send the message to the chatroom
				chatroom.MessageChannel <- msg

				// Respond with success
				writer.WriteHeader(http.StatusCreated)
				writer.Header().Set(string(common.ContentType), string(common.ApplicationJson))
				response := models.NewMessageResponse(common.Success)

				json.NewEncoder(writer).Encode(response)

				return
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
