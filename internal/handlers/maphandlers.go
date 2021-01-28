package handlers

import (
	"fmt"
	"net/http"

	"../core"
	"../database"
)

// NewShortenHandler creates a handler which redirects from a shortened URL
func NewMapHandler(log core.Logger, db database.Database) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		log.Infof("Attempting redirect on %s\n", r.URL.Path)
		redirect, err := db.Get(r.URL.Path[1:])
		if err == nil {
			http.Redirect(rw, r, redirect, http.StatusFound)
			return
		}

		// If no redirect, fall back to fallback
		log.Infof("Unable to find URL for %s\n", r.URL.Path)
		rw.Write([]byte(fmt.Sprintf("No redirect registered for %s\n", r.URL.Path)))
	}

	return http.HandlerFunc(fn)
}
