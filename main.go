package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! This is a simple Go web server.")
	})

	// Start the server on port 8080
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
