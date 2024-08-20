package models

import (
	"slices"
	"sync"
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
	Users          []User
	MessageChannel chan Message
	messages       []Message
	Node           *common.Node
	RwMutex        sync.RWMutex
}

func NewChatroom(node *common.Node) *Chatroom {
	chatroom := Chatroom{
		Users:          []User{},
		MessageChannel: make(chan Message, MESSAGE_LIMIT),
		messages:       []Message{},
		Node:           node,
	}

	go chatroom.handleMessages()
	return &chatroom
}

// Handles message transition between
func (chatroom *Chatroom) handleMessages() {
	for msg := range chatroom.MessageChannel {
		select {
		case <-chatroom.Node.Context.Done():
			close(chatroom.MessageChannel)
			return

		default:
			{
				chatroom.RwMutex.Lock()
				// Delete the last message
				if len(chatroom.messages) >= MESSAGE_LIMIT {
					chatroom.messages = slices.Delete(chatroom.messages, 0, 1)
				}
				chatroom.messages = append(chatroom.messages, msg)
				chatroom.RwMutex.Unlock()
			}
		}
	}
}

func (chatroom *Chatroom) GetMessages() []Message {
	chatroom.RwMutex.RLock()
	messages := chatroom.messages
	chatroom.RwMutex.RUnlock()
	return messages
}
