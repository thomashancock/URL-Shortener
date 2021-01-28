package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thomashancock/URL-Shortener/internal/utils"
	"github.com/thomashancock/URL-Shortener/internal/database"

	"github.com/stretchr/testify/assert"
)

// Test_ShortenHandler_EmptyDB tests ShortenHandler when the db is empty
func Test_ShortenHandler_EmptyDB(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	sh := NewShortenHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/shorten?url=Test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	sh.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Test can now be accessed on /0\n"
	assert.Equal(t, rr.Body.String(), expected)

	expectedLog := "Test can now be accessed on /0\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}

// Test_ShortenHandler_EntryExists tests ShortenHandler when the db already has entries
func Test_ShortenHandler_EntryExists(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	testDB.Add("0", "TestEntry")

	sh := NewShortenHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/shorten?url=Test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	sh.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Test can now be accessed on /1\n"
	assert.Equal(t, rr.Body.String(), expected)

	expectedLog := "Test can now be accessed on /1\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}
