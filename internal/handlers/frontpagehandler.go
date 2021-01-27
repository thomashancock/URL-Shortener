package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func FrontpageHandler(rw http.ResponseWriter, r *http.Request) {
	log.Infoln("Responding with Front Page")
	rw.Write([]byte("Thomas' URL Shortener!\n"))
	rw.Write([]byte("go to /shorten to add your URL and get a shortened version\n"))
	rw.Write([]byte("syntax: /shorten?url=<your url here>\n"))
}
