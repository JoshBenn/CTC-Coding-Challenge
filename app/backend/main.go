package main

import (
	"fmt"
	"net/http"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
)

func main() {
	fmt.Println("Initializing backend")

	node := common.NewNode()
	node.Output <- "Backend Initialized\n"
	defer node.Shutdown()

	mux := http.NewServeMux()
	// To get go to shut up about unused
	fmt.Println(mux)

}
