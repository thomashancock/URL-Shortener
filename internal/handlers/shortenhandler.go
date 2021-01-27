package handlers

import (
	"fmt"
	"net/http"

	"../core"
	"../database"
	"../utils"
)

func NewShortenHandler(log core.Logger, db database.Database) http.HandlerFunc {
	// Declare a URL generator
	urlGenerator := utils.NewURLGenerator()

	// Handler function which adds the url to the redirects map
	fn := func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Query()["url"][0]

		log.Infof("Attempting shortening for %s\n", url)

		shortURL, err := urlGenerator()
		if err != nil {
			rw.Write([]byte(fmt.Sprintf("Unable to provide short URL for %s\n", url)))
			log.Errorf("Unable to provide short URL for %s\n", url)
			return
		}

		err = db.Add(shortURL, url)
		if err != nil {
			rw.Write([]byte(fmt.Sprintf("Unable to provide short URL for %s\n", url)))
			log.Errorf("Unable to provide short URL for %s\n", url)
			return
		}
		rw.Write([]byte(fmt.Sprintf("%s can now be accessed on /%s\n", url, shortURL)))
		log.Infof("%s can now be accessed on /%s\n", url, shortURL)
	}

	return http.HandlerFunc(fn)
}
