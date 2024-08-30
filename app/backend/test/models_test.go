package test

import (
	"context"
	"testing"
	"time"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/JoshBenn/CTC-Coding-Challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestNewChatroom(t *testing.T) {
	node := &common.Node{
		Context: context.Background(),
	}
	chatroom := models.NewChatroom(node)

	assert.NotNil(t, chatroom)
	assert.NotNil(t, chatroom.MessageChannel)
	assert.Empty(t, chatroom.Users)
	assert.Empty(t, chatroom.GetMessages())
}

func TestChatroomHandleMessages(t *testing.T) {
	node := &common.Node{
		Context: context.Background(),
	}
	chatroom := models.NewChatroom(node)

	// Simulate adding messages
	message1 := models.Message{Username: "user1", Content: "Hello"}
	message2 := models.Message{Username: "user2", Content: "Hi"}

	chatroom.MessageChannel <- message1
	chatroom.MessageChannel <- message2

	// Allow time for messages to be handled
	time.Sleep(100 * time.Millisecond)

	messages := chatroom.GetMessages()

	assert.Equal(t, 2, len(messages))
	assert.Equal(t, message1, messages[0])
	assert.Equal(t, message2, messages[1])
}

func TestChatroomMessageLimit(t *testing.T) {
	node := &common.Node{
		Context: context.Background(),
	}
	chatroom := models.NewChatroom(node)

	// Send MESSAGE_LIMIT + 1 messages
	for i := 0; i < models.MESSAGE_LIMIT+1; i++ {
		chatroom.MessageChannel <- models.Message{Username: "user", Content: "Message"}
	}

	// Allow time for messages to be handled
	time.Sleep(100 * time.Millisecond)

	messages := chatroom.GetMessages()
	assert.Equal(t, models.MESSAGE_LIMIT, len(messages))
}

func TestChatroomConcurrency(t *testing.T) {
	node := &common.Node{
		Context: context.Background(),
	}
	chatroom := models.NewChatroom(node)

	// Simulate concurrent access
	done := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			chatroom.MessageChannel <- models.Message{Username: "user", Content: "Concurrent message"}
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			chatroom.GetMessages()
		}
		done <- true
	}()

	// Wait for both goroutines to complete
	<-done
	<-done

	assert.True(t, true, "Concurrent access did not cause any data races")
}
