package handlers

import (
	"net/http"

	"../core"
)

func NewFrontpageHandler(log core.Logger) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Thomas' URL Shortener!\n"))
		rw.Write([]byte("go to /shorten to add your URL and get a shortened version\n"))
		rw.Write([]byte("syntax: /shorten?url=<your url here>\n"))
	}
	return http.HandlerFunc(fn)
}
