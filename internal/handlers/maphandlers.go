package handlers

import (
	"net/http"

	"../database"

	log "github.com/sirupsen/logrus"
)

func MapHandler(db database.Database, fallback http.Handler) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		log.Infof("Attempting redirect on %s\n", r.URL.Path)
		redirect, err := db.Get(r.URL.Path[1:])
		if err == nil {
			http.Redirect(rw, r, redirect, http.StatusFound)
			return
		}

		// If no redirect, fall back to fallback
		fallback.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}
