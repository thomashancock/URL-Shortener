package main

import (
	"net/http"
	"fmt"
)


func FrontpageHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Thomas' URL Shortener!")
	fmt.Fprintln(rw, "go to /shorten to add your URL and get a shortened version")
	fmt.Fprintln(rw, "syntax: /shorten?url=<your url here>")
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


func ShortenHandler(redirects map[string]string) http.HandlerFunc {
	// Declare a URL generator
	urlGenerator := URLGenerator()

	// Handler function which adds the url to the redirects map
	fn := func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Query()["url"][0]
		shortURL := "/" + urlGenerator()
		redirects[shortURL] = url
		fmt.Fprintln(rw, url, "can now be accessed on", shortURL)
	}

	return http.HandlerFunc(fn)
}
