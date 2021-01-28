package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../utils"
	"../database"

	"github.com/stretchr/testify/assert"
)

// Test_RouteHandler_Frontpage tests requests for /
func Test_RouteHandler_Frontpage(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	router := NewRouteHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Thomas' URL Shortener!\ngo to /shorten to add your URL and get a shortened version\nsyntax: /shorten?url=<your url here>\n"
	assert.Equal(t, rr.Body.String(), expected)
}

// Test_RouteHandler_Frontpage tests requests for /shorten
func Test_RouteHandler_Shorten(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	router := NewRouteHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/shorten?url=Test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Test can now be accessed on /0\n"
	assert.Equal(t, rr.Body.String(), expected)

	expectedLog := "Test can now be accessed on /0\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}

// Test_RouteHandler_Other tests requests for other paths
func Test_RouteHandler_Other(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	testDB.Add("0", "TestEntry")

	router := NewRouteHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusFound)

	expected := "<a href=\"/TestEntry\">Found</a>.\n\n"
	assert.Equal(t, rr.Body.String(), expected)

	expectedLog := "Attempting redirect on /0\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}
