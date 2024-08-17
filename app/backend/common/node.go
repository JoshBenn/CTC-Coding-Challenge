package common

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const logFile = "LogFile"

type node struct {
	File    *os.File
	Logger  *slog.Logger
	Log     chan Log
	Output  chan string
	Context context.Context
	Cancel  context.CancelFunc
	Server  *http.Server
}

func NewNode() *node {
	file, err := os.OpenFile(os.Getenv(string(logFile)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(file, nil))
	slog.SetDefault(logger)

	context, cancel := context.WithCancel(context.Background())
	n := node{
		File:    file,
		Logger:  logger,
		Log:     make(chan Log, 100),
		Output:  make(chan string, 100),
		Context: context,
		Cancel:  cancel,
	}

	n.init()
	return &n
}

func (n *node) init() {
	go n.handleOutput()
	go n.handleLog()
}

func (n *node) handleOutput() {
	for output := range n.Output {
		select {
		case <-n.Context.Done():
			return

		default:
			fmt.Println(output)
		}
	}
}

func (n *node) handleLog() {
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

func (n *node) Shutdown() {
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
