package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// ParseForm is necessary to extract form data from the request.
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusInternalServerError)
		return
	}

	// Handle POST request
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintln(w, "POST request successful")
		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Address: %s\n", address)
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r) // Use http.NotFound to handle 404 responses.
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello")
}

func main() {
	// Serving static files from the "static" directory.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Handling form submissions and hello route.
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	// Start the HTTP server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
