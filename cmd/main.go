package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"../internal/handlers"
)


func main() {
	// Declare server
	mux := http.NewServeMux()

	// Add handler for main page
	mux.HandleFunc("/", handlers.FrontpageHandler)

	// Create map to store URLs
	redirects := make(map[string]string)

	// Handler to generate and register shortened URLs
	sh := handlers.ShortenHandler(redirects)
	mux.HandleFunc("/shorten", sh)

	// Handler to redirect URLs
	mh := handlers.MapHandler(redirects, mux)

	// Run server
	go http.ListenAndServe(":8080", mh)
	log.Println("Listening on http://localhost:8080")

	// Run until interrupt signal received
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
