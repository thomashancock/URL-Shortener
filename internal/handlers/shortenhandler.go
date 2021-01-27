package handlers

import (
	"net/http"
	"fmt"

	"../utils"
)


func ShortenHandler(redirects map[string]string) http.HandlerFunc {
	// Declare a URL generator
	urlGenerator := utils.URLGenerator()

	// Handler function which adds the url to the redirects map
	fn := func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Query()["url"][0]
		shortURL := "/" + urlGenerator()
		redirects[shortURL] = url
		fmt.Fprintln(rw, url, "can now be accessed on", shortURL)
	}

	return http.HandlerFunc(fn)
}
