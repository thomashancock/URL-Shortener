package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"fmt"
	"strconv"
)


func GetShortURL(val int) string {
	// For now, just turn the integer into a string
	out := strconv.Itoa(val)

	return out
}


func MapHandler(redirects map[string]string, fallback http.Handler) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if redirects[path] != "" {
			http.Redirect(rw, r, redirects[path], http.StatusFound)
		}

		// If no redirect, fall back to fallback
		fallback.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}


func FrontpageHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Thomas' URL Shortener!")
	fmt.Fprintln(rw, "go to /shorten to add your URL and get a shortened version")
	fmt.Fprintln(rw, "syntax: /shorten?url=<your url here>")
}


func ShortenHandler(redirects map[string]string) http.HandlerFunc {
	// count used to give each url a unique redirect url
	count := 0

	fn := func(rw http.ResponseWriter, r *http.Request) {
		shortURL := "/" + GetShortURL(count)
		count += 1

		url := r.URL.Query()["url"][0]
		redirects[shortURL] = r.URL.Query()["url"][0]
		fmt.Fprintln(rw, url, "can now be accessed on", shortURL)
	}

	return http.HandlerFunc(fn)
}


func main() {
	// Declare server
	mux := http.NewServeMux()

	// Add handler for main page
	mux.HandleFunc("/", FrontpageHandler)

	// Create map to store URLs
	redirects := make(map[string]string)

	// Handler to generate and register shortened URLs
	sh := ShortenHandler(redirects)
	mux.HandleFunc("/shorten", sh)

	// Handler to redirect URLs
	mh := MapHandler(redirects, mux)

	// Run server
	go http.ListenAndServe(":8080", mh)
	log.Println("Listening on http://localhost:8080")

	// Watch for interrupt signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
