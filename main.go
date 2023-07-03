package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler) // Register the handler function for the root route
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!") // Send "Hello, World!" as the response
}
