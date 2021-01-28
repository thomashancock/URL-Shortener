package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../utils"

	"github.com/stretchr/testify/assert"
)

// Test_MapHandler_EntryExists tests MapHandler when the entry is in the database
func Test_MapHandler_EntryExists(t *testing.T) {
	testLog := &utils.TestLogger{}
	testDB := utils.NewTestDatabase(testLog)

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
}
