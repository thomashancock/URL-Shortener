package handlers

import (
	"net/http"
	"fmt"

	"../utils"
	"../database"
)


func ShortenHandler(db database.Database) http.HandlerFunc {
	// Declare a URL generator
	urlGenerator := utils.NewURLGenerator()

	// Handler function which adds the url to the redirects map
	fn := func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Query()["url"][0]

		shortURL, err := urlGenerator()
		if err != nil {
			rw.Write([]byte(fmt.Sprintf("Unable to provide short URL for %s", url)))
			return
		}

		err = db.Add(shortURL, url)
		if err != nil {
			rw.Write([]byte(fmt.Sprintf("Unable to provide short URL for %s", url)))
			return
		}
		rw.Write([]byte(fmt.Sprintf("%s can now be accessed on /%s", url, shortURL)))
	}

	return http.HandlerFunc(fn)
}
