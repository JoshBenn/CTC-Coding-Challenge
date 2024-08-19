package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/JoshBenn/CTC-Coding-Challenge/controllers"
)

func main() {
	fmt.Println("Initializing backend")

	// Serve the frontend
	http.Handle("/", http.FileServer(http.Dir("../frontend/dist")))

	node := common.NewNode()
	node.Output <- "Backend Initialized\n"
	defer node.Shutdown()

	mux := http.NewServeMux()
	// To get go to shut up about unused
	mux.HandleFunc(string(common.Authentication), controllers.AuthenticationHandler(node))
	mux.HandleFunc(string(common.Chat), controllers.ChatHandler(node))

	server := &http.Server{
		Addr:    os.Getenv(string(common.BackendAddress)),
		Handler: mux,
	}
	node.Server = server

	go func() {
		if err := http.ListenAndServe(os.Getenv(string(common.BackendAddress)), mux); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starrting on server: ", err)
			node.Cancel()
		}
	}()

	<-node.Context.Done()
}
