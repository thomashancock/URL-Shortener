package handlers

import (
	"net/http"

	"../core"
	"../database"
)

func NewMapHandler(log core.Logger, db database.Database, fallback http.Handler) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		log.Infof("Attempting redirect on %s\n", r.URL.Path)
		redirect, err := db.Get(r.URL.Path[1:])
		if err == nil {
			http.Redirect(rw, r, redirect, http.StatusFound)
			return
		}

		// If no redirect, fall back to fallback
		log.Errorf("Unable to find URL for %s\n", r.URL.Path)
		fallback.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}
