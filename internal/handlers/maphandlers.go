package handlers

import (
	"net/http"

	"../database"
)


func MapHandler(db database.Database, fallback http.Handler) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		redirect, err := db.Get(r.URL.Path[1:])
		if err == nil {
			http.Redirect(rw, r, redirect, http.StatusFound)
		}

		// If no redirect, fall back to fallback
		fallback.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}
