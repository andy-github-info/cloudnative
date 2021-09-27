package main

import (
	"io"
	"log"
	"net/http"
)

func newPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This is the people handler.")
	})
}

func main() {
	mux := http.NewServeMux()

	// Create sample handler to returns 404
	mux.Handle("/resources", http.NotFoundHandler())

	// Create sample handler that returns 200
	mux.Handle("/resources/people/", newPeopleHandler())

	log.Fatal(http.ListenAndServe(":80", mux))
}