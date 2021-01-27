package handlers

import (
	"net/http"
)


func MapHandler(redirects map[string]string, fallback http.Handler) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if redirects[path] != "" {
			http.Redirect(rw, r, redirects[path], http.StatusFound)
		}

		// If no redirect, fall back to fallback
		fallback.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}
