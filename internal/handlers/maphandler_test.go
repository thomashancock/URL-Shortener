package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../utils"
	"../database"

	"github.com/stretchr/testify/assert"
)

// Test_MapHandler_EntryExists tests MapHandler when the entry is in the database
func Test_MapHandler_EntryExists(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	testDB.Add("0", "TestEntry")

	mh := NewMapHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mh.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusFound)

	expected := "<a href=\"/TestEntry\">Found</a>.\n\n"
	assert.Equal(t, rr.Body.String(), expected)

	expectedLog := "Attempting redirect on /0\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}

// Test_MapHandler_NoEntry tests MapHandler when their is no entry in the database
func Test_MapHandler_NoEntry(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := database.NewTestDatabase(testLog)

	mh := NewMapHandler(testLog, testDB)

	req, err := http.NewRequest("GET", "/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mh.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "No redirect registered for /0\n"
	assert.Equal(t, rr.Body.String(), expected)

	expectedLog := "Unable to find URL for /0\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}
