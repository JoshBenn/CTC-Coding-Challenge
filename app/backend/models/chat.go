package models

import "time"

// Represents a single message in the application
type Message struct {
	Username string
	Time     time.Time
	Content  string
}

// Represents a single chatroom
// -- In prod, will have more than a single chatroom available, but for the purposes of this exercise will just maintain a single one
type Chatroom struct {
	Users    []User
	Messages []Message
}
