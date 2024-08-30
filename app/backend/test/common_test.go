package test

import (
	"os"
	"testing"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/stretchr/testify/assert"
)

// TestNewNode tests the initialization of a Node
func TestNewNode(t *testing.T) {
	// Set up environment variables
	os.Setenv("LOG_FILE", "test.log")
	defer os.Unsetenv("LOG_FILE")

	// Create a new Node
	node := common.NewNode()
	defer node.Shutdown()

	assert.NotNil(t, node)
	assert.NotNil(t, node.File)
	assert.NotNil(t, node.Logger)
	assert.NotNil(t, node.Log)
	assert.NotNil(t, node.Output)
	assert.NotNil(t, node.Context)
	assert.NotNil(t, node.Cancel)
}

// TestNodeShutdown tests the shutdown behavior of the Node
func TestNodeShutdown(t *testing.T) {
	os.Setenv("LOG_FILE", "test.log")
	defer os.Unsetenv("LOG_FILE")

	node := common.NewNode()

	node.Shutdown()
	assert.True(t, node.Context.Err() != nil)

	// Check if channels are closed
	_, outputOk := <-node.Output
	_, logOk := <-node.Log

	assert.False(t, outputOk)
	assert.False(t, logOk)
}

// Test new log
func TestNewLog(t *testing.T) {
	testDebug := common.NewLog(common.Debug, "test")
	testWarn := common.NewLog(common.Warn, "test")
	testInfo := common.NewLog(common.Info, "test")
	testError := common.NewLog(common.Error, "test")

	assert.Equal(t, testDebug, common.Log{
		Level:   common.Debug,
		Message: "test",
	})
	assert.Equal(t, testWarn, common.Log{
		Level:   common.Warn,
		Message: "test",
	})
	assert.Equal(t, testInfo, common.Log{
		Level:   common.Info,
		Message: "test",
	})
	assert.Equal(t, testError, common.Log{
		Level:   common.Error,
		Message: "test",
	})
}
