package common

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

// Custom errors for backend
type NodeError string

const (
	ServerCompilationError  NodeError = "Server compilation error"
	InternalServiceError    NodeError = "Internal service error"
	MissingJwtSecretError   NodeError = "Missing JWT secret"
	DatabaseConnectionError NodeError = "Could not generate a connection to the database"
)

// Contains the backend primary structure
type Node struct {
	File    *os.File
	Logger  *slog.Logger
	Log     chan Log
	Output  chan string
	Context context.Context
	Cancel  context.CancelFunc
	Server  *http.Server
}

// Generates and returns a new node
func NewNode() *Node {
	file, err := os.OpenFile(os.Getenv(string(logFile)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, permissions)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(file, nil))
	slog.SetDefault(logger)

	context, cancel := context.WithCancel(context.Background())
	n := Node{
		File:    file,
		Logger:  logger,
		Log:     make(chan Log, 100),
		Output:  make(chan string, 100),
		Context: context,
		Cancel:  cancel,
	}

	n.init()
	n.Output <- fmt.Sprintf("Node started at %v", time.Now().UTC())
	return &n
}

// Initializes the I/O of the backend
func (n *Node) init() {
	go n.handleOutput()
	go n.handleLog()
	go n.handleInput()
}

// Handles all input - specifies looking for exit to shut down the backend
func (n *Node) handleInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		select {
		case <-n.Context.Done():
			return

		default:
			text := scanner.Text()
			switch strings.ToLower(text) {
			case "exit":
				n.Cancel()
				return
			}
		}
	}
}

// Handles all output through the backend output channel
func (n *Node) handleOutput() {
	for output := range n.Output {
		select {
		case <-n.Context.Done():
			return

		default:
			fmt.Println(output)
		}
	}
}

// Handles all logging for the system
func (n *Node) handleLog() {
	for log := range n.Log {
		select {
		case <-n.Context.Done():
			return

		default:
			{
				switch log.Level {
				case Debug:
					slog.Debug(log.Message)
				case Info:
					slog.Info(log.Message)
				case Warn:
					slog.Warn(log.Message)
				case Error:
					slog.Error(log.Message)
				default:
					n.Output <- "Invalid log level passed"
				}
			}
		}
	}
}

// Shuts down the backend, ensuring all buffers are emptied and closed
func (n *Node) Shutdown() {
	fmt.Println("Shutting down server")

	n.Cancel()
	if n.Server != nil {
		context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := n.Server.Shutdown(context); err != nil {
			fmt.Println("Error shutting down server:", err)
		}
	}

	// Close the buffers
	close(n.Output)
	close(n.Log)

	fmt.Println("Emptying Buffers:")
	// Empty the buffers
	for output := range n.Output {
		fmt.Println("-", output)
	}

	for log := range n.Log {
		switch log.Level {
		case Debug:
			slog.Debug(log.Message)
		case Info:
			slog.Info(log.Message)
		case Warn:
			slog.Warn(log.Message)
		case Error:
			slog.Error(log.Message)
		default:
			n.Output <- "Invalid log level passed"
		}
	}

	// Close the file
	n.File.Close()
}
