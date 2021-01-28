package handlers

import (
	"net/http"

	"../core"
	"../database"
)

// NewRouteHandler calls the correct handler based on the passed path
func NewRouteHandler(log core.Logger, db database.Database) http.HandlerFunc {

	fh := NewFrontpageHandler(log)
	sh := NewShortenHandler(log, db)
	mh := NewMapHandler(log, db, fh)

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
