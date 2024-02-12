package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Handler for the root "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	// add one comment

	// Handler for "/submit"
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Print the request body
		fmt.Printf("Received submission: %s\n", string(body))

		// Respond to the client
		fmt.Fprintf(w, "Received submission: %s", body)
	})

	// Start the server
	log.Println("Starting HTTP server on http://localhost:8080")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
