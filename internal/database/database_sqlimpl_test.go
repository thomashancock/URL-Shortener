package database

import (
	"testing"
	"os"

	"../utils"

	"github.com/stretchr/testify/assert"
)

// Test_Database tests the basic functionality of the database
func Test_Database(t *testing.T) {
	testLog := &utils.TestLogger{}
	dbFile := "sqlite-test.db"

	// Remove file if it already exists
	if fileExists(dbFile) {
		err := os.Remove(dbFile)
		if err != nil {
			t.Fatal(err)
		}
	}

	testDB := NewSQLDatabase(testLog, dbFile)
	expectedLog := "Created SQL DB\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)

	nEntries, err := testDB.NEntries()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nEntries, 0)

	// Add entry to DB
	testDB.Add("0", "Test")

	nEntries, err = testDB.NEntries()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nEntries, 1)

	// Retreive entry from DB
	entry, err := testDB.Get("0")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Test"
	assert.Equal(t, entry, expected)
}
