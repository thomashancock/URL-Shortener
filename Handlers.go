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
