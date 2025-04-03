package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// The Header().Add() method can be used to add headers to the
	// HTTP response header map. Param 1 => name, Param 2 => value
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))
}

// Another handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Wildcards can be extracted using r.PathValue() method
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Yet another handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Another handler!
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// The w.WriteHeader() method can be used to set the HTTP response code
	// Note: This function only functions properly once -- any subsequent calls will do nothing, other than log a warning
	w.WriteHeader(http.StatusCreated) // http.StatusCreated == 201

	w.Write([]byte("Save a new snippet..."))
}

func main() {
	port := ":4000"

	// Use the http.NewServeMux() function to initialize a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on ", port)

	// Use the http.ListenAndServe() function to start a new web server. Two parameters:
	// 1.) The tcp network address to listen on
	// 2.) The servemux
	// If http.ListenAndServe() returns an error, we use the log.Fatal() function to log the error
	// message and exit. Note that http.ListenAndServe errors are always non-nil.
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
