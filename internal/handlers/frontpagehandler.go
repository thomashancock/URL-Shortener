package handlers

import (
	"net/http"
	"fmt"
)


func FrontpageHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Thomas' URL Shortener!")
	fmt.Fprintln(rw, "go to /shorten to add your URL and get a shortened version")
	fmt.Fprintln(rw, "syntax: /shorten?url=<your url here>")
}
