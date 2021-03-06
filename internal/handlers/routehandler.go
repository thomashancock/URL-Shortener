package handlers

import (
	"net/http"

	"github.com/thomashancock/URL-Shortener/internal/core"
	"github.com/thomashancock/URL-Shortener/internal/database"
)

// NewRouteHandler creates a handler which calls the another handler based on the passed path
func NewRouteHandler(log core.Logger, db database.Database) http.HandlerFunc {

	fh := NewFrontpageHandler(log)
	sh := NewShortenHandler(log, db)
	mh := NewMapHandler(log, db)

	router := func(rw http.ResponseWriter, r *http.Request) {
		log.Infof("Routing HTTP request: %s\n", r.URL.Path)

		switch r.URL.Path {
		case "/":
			fh(rw, r)
		case "/shorten":
			sh(rw, r)
		default:
			mh(rw, r)
		}
	}

	return http.HandlerFunc(router)
}
