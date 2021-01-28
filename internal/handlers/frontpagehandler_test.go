package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thomashancock/URL-Shortener/internal/utils"

	"github.com/stretchr/testify/assert"
)

// Test_FrontpageHandler tests FrontpageHandler
func Test_FrontpageHandler(t *testing.T) {
	testLog := &utils.TestLogger{}
	fh := NewFrontpageHandler(testLog)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fh.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Thomas' URL Shortener!\ngo to /shorten to add your URL and get a shortened version\nsyntax: /shorten?url=<your url here>\n"
	assert.Equal(t, rr.Body.String(), expected)
}
