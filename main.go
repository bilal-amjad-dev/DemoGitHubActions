package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! This is test commit.")
}

func main() {
	http.HandleFunc("/", handler) // Handle requests to the root URL
	port := "8000"                // Define the port to listen on
	fmt.Printf("Go web app listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil)) // Start the web server
}
