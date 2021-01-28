package handlers

import (
	"fmt"
	"net/http"

	"github.com/thomashancock/URL-Shortener/internal/core"
	"github.com/thomashancock/URL-Shortener/internal/database"
	"github.com/thomashancock/URL-Shortener/internal/utils"
)

// NewShortenHandler creates a handler which registers a shortened URL
// shortened URL generation is handled by utils.urlgenerator
func NewShortenHandler(log core.Logger, db database.Database) http.HandlerFunc {
	// Declare a URL generator
	nURLs, err := db.NEntries()
	if err != nil {
		log.Fatalf("Unable to get number of entries in db")
	}
	urlGenerator := utils.NewURLGenerator(log, nURLs)

	// Handler function which adds the url to the redirects map
	fn := func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Query()["url"][0]

		log.Infof("Attempting shortening for %s\n", url)

		shortURL, err := urlGenerator.Get()
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
