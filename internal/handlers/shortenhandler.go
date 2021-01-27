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
		rw.Write([]byte(fmt.Sprintf("%s can now be accessed on %s", url, shortURL)))
	}

	return http.HandlerFunc(fn)
}
