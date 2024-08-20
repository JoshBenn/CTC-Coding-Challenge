package models

import (
	"slices"
	"time"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
)

const MESSAGE_LIMIT = 1000

// Represents a single message in the application
type Message struct {
	Username string
	Time     time.Time
	Content  string
}

// Represents a single chatroom
// -- In prod, will have more than a single chatroom available, but for the purposes of this exercise will just maintain a single one
type Chatroom struct {
	Users       []User
	MessageChan chan Message
	Messages    []Message
	Node        *common.Node
}

func NewChatroom(node *common.Node) *Chatroom {
	chatroom := Chatroom{
		Users:       []User{},
		MessageChan: make(chan Message, MESSAGE_LIMIT),
		Messages:    []Message{},
		Node:        node,
	}

	go chatroom.handleMessages()
	return &chatroom
}

func (chatroom *Chatroom) handleMessages() {
	for msg := range chatroom.MessageChan {
		select {
		case <-chatroom.Node.Context.Done():
			return

		default:
			{
				// Delete the last message
				if len(chatroom.Messages) >= MESSAGE_LIMIT {
					chatroom.Messages = slices.Delete(chatroom.Messages, 0, 1)
				}
				chatroom.Messages = append(chatroom.Messages, msg)

			}
		}
	}
}
